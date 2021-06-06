package main

import (
	"fmt"
	"io/ioutil"
)

func Search(path string) func(query string) func(log bool) (*bool, *int) {
	matches := 0

	finderCount := 1
	foundMatch := make(chan bool)
	findDone := make(chan bool)
	searchRequest := make(chan string)
	ok := false

	return func(query string) func(log bool) (*bool, *int) {
		search := func(path string, log bool) {
			files, err := ioutil.ReadDir(path)
			if err == nil {
				for _, file := range files {
					name := file.Name()
					if file.IsDir() {
						searchRequest <- path + name + "/"
					} else {
						if name == query {
							if log {
								fmt.Println(path + name)
							}
							foundMatch <- true
						}
					}
				}
			}
			findDone <- true
		}
		waitForSearch := func(log bool) {
			for {
				select {
				case path := <-searchRequest:
					finderCount++
					go search(path, log)
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
		return func(log bool) (*bool, *int) {
			go search(path, log)
			go waitForSearch(log)

			return &ok, &matches
		}
	}
}
