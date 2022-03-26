// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type urlInfo struct {
	links []string
	depth int
}

var tokens = make(chan struct{}, 10)

func crawl(url string, depth int) *urlInfo {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("depth %d %s\n", depth, url)
	return &urlInfo{list, depth + 1}
}

var depthLimit = flag.Int("depth", 1, "depth of crawl\n")

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Invalid number of Argument\n")
	}
	flag.Parse()

	worklist := make(chan urlInfo)
	var n int // number of pending sends to worklist
	// Start with the command-line arguments.
	n++
	go func() { worklist <- urlInfo{os.Args[2:], 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				if list.depth > *depthLimit {
					break
				}
				n++
				go func(link string, depth int) {
					worklist <- *crawl(link, depth)
				}(link, list.depth)
			}
		}
	}
	fmt.Printf("Finish!!\n")
}
