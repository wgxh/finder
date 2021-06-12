package main

import (
	"flag"
	"fmt"
	"time"
)

var path = flag.String("p", "/home/wgxh-cli/Code/", "")
var query = flag.String("q", "main.js", "")
var logMode = flag.String("log", "no-log", "")

func main() {
	start := time.Now()

	flag.Parse()
	ok, matches := Search(*path)(*query)(*logMode)
	for {
		if *ok {
			break
		}
	}
	fmt.Println("Found", *matches, "matches.")
	fmt.Println("take", time.Since(start))
}
