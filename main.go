package main

import (
	"os"
)

func main() {
	finder := Finder{
		path:          os.Args[1],
		query:         os.Args[2],
		matches:       0,
		finderCount:   1,
		finderDone:    make(chan bool),
		foundMatch:    make(chan bool),
		searchRequest: make(chan string),
	}
	finder.search()
}
