// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"ch5/ex5_13/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, v := range list {
		if strings.HasPrefix(v, url) {
			resp, err := http.Get(v)
			if err != nil {
				log.Print(err)
				continue
			}
			if resp.StatusCode != http.StatusOK {
				log.Printf("getting %s: %s", v, resp.Status)
				resp.Body.Close()
				continue
			}
			b, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				log.Printf("reading %s: %v", v, err)
				continue
			}
			urlSplitDir := strings.Split(v, url)
			slashSplitDir := strings.Split(strings.Join(urlSplitDir, ""), "/")
			dirName := strings.Split(url, "https://")[1]
			dirName = strings.Split(dirName, "/")[0]
			dirPath := dirName
			for _, dir := range slashSplitDir {
				if strings.HasSuffix(dir, ".html") {
					break
				}
				dirPath += dir
				if !strings.HasSuffix(dirPath, "/") {
					dirPath += "/"
				}
				if _, err = os.Stat(dirPath); os.IsNotExist(err) {
					err = os.Mkdir(dirPath, 0644)
					if err != nil {
						log.Printf("making %s: %v", dirPath, err)
						continue
					}
					err = os.Chmod(dirPath, 0744)
					if err != nil {
						log.Printf("changing permmision of %s: %v", dirPath, err)
						continue
					}
				}
			}
			var filepath string
			if !strings.HasSuffix(filepath, ".html") {
				filepath = dirPath + "index.html"
			}
			err = os.WriteFile(filepath, b, 0644)
			if err != nil {
				log.Printf("writing %s: %v", filepath, err)
				continue
			}
		}
	}
	return list
}

//!-crawl

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
