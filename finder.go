package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Finder struct {
	path       string
	query      string
	matches    int
	activePath string

	finderCount   int
	searchRequest chan string
	foundMatch    chan bool
	finderDone    chan bool
}

func (finder *Finder) searchMain(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if file.IsDir() {
				fmt.Println(name)
				finder.searchRequest <- path + name + "/"
			} else {
				if name == finder.query {
					finder.foundMatch <- true
				}
			}
		}
		if master {
			finder.finderDone <- true
		}
	} else {
		fmt.Println(err)
	}
}

func (finder *Finder) waitForFinders() {
	for {
		select {
		case path := <-finder.searchRequest:
			finder.finderCount++
			finder.activePath = path
			fmt.Println(path)
			go finder.searchMain(path, true)
		case <-finder.foundMatch:
			finder.matches++
		case <-finder.finderDone:
			finder.finderCount--
			if finder.finderCount == 0 {
				return
			}
		default:
			continue
		}
	}
}

func (finder *Finder) search() {
	fmt.Println(finder.path)
	fmt.Println(finder.query)
	start := time.Now()
	go finder.searchMain(finder.path, true)
	finder.waitForFinders()
	fmt.Println("Found", finder.matches, "matches")
	fmt.Println("take", time.Since(start))
}

func (finder *Finder) setPath(path string) {
	finder.path = path
}

func (finder *Finder) setQuery(query string) {
	finder.query = query
}
