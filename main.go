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
	isLog := flag.Bool("log", false, "")

	flag.Parse()
	ok, matches := Search(*path)(*query)(*isLog)
	for {
		if *ok {
			break
		}
	}
	fmt.Println("Found", *matches, "matches.")
	fmt.Println("take", time.Since(start))
}
