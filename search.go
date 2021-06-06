package main

import (
	"io/ioutil"
)

func Search(path string) func(query string) func() (*bool, *int) {
	matches := 0

	finderCount := 1
	foundMatch := make(chan bool)
	findDone := make(chan bool)
	searchRequest := make(chan string)
	ok := false

	return func(query string) func() (*bool, *int) {
		search := func(path string) {
			files, err := ioutil.ReadDir(path)
			if err == nil {
				for _, file := range files {
					name := file.Name()
					if file.IsDir() {
						searchRequest <- path + name + "/"
					} else {
						if name == query {
							foundMatch <- true
						}
					}
				}
			}
			findDone <- true
		}
		waitForSearch := func() {
			for {
				select {
				case path := <-searchRequest:
					finderCount++
					go search(path)
				case <-foundMatch:
					matches++
				case <-findDone:
					finderCount--
					if finderCount == 0 {
						ok = true
						return
					}
				}
			}
		}
		return func() (*bool, *int) {
			go search(path)
			go waitForSearch()

			return &ok, &matches
		}
	}
}
