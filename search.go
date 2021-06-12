package main

import (
	"fmt"
	"io/ioutil"
)

func Search(path string) func(query string) func(logMode string) (*bool, *int) {
	matches := 0

	finderCount := 1
	foundMatch := make(chan bool)
	findDone := make(chan bool)
	searchRequest := make(chan string)
	ok := false

	return func(query string) func(logMode string) (*bool, *int) {
		search := func(path string, logMode string) {
			files, err := ioutil.ReadDir(path)
			if err == nil {
				for _, file := range files {
					name := file.Name()
					if file.IsDir() {
						if logMode == "search" {
							fmt.Println(path + name + "/")
						}
						searchRequest <- path + name + "/"
					} else {
						if name == query {
							if logMode == "match" {
								fmt.Println(path + name)
							}
							foundMatch <- true
						}
					}
				}
			}
			findDone <- true
		}
		searchEffect := func(logMode string) {
			for {
				select {
				case path := <-searchRequest:
					finderCount++
					go search(path, logMode)
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
		return func(logMode string) (*bool, *int) {
			go search(path, logMode)
			go searchEffect(logMode)

			return &ok, &matches
		}
	}
}
