package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	path := flag.String("p", "/home/wgxh-cli/Code/", "")
	query := flag.String("q", "main.js", "")
	flag.Parse()
	ok, matches := Search(*path)(*query)()
	for {
		if *ok {
			break
		}
	}
	fmt.Println("Found", *matches, "matches.")
	fmt.Println("take", time.Since(start))
}
