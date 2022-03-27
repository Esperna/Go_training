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
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

type urlInfo struct {
	links []string
	depth int
}

var tokens = make(chan struct{}, 10)

func crawl(urlStr string, depth int) *urlInfo {
	tokens <- struct{}{}
	list, err := links.Extract(urlStr)
	<-tokens
	if err != nil {
		log.Printf("failed to extrackt %s: %s", urlStr, err)
	}
	for _, item := range list {
		if strings.HasPrefix(item, urlStr) {
			extractUrl, err := url.Parse(item)
			if err != nil {
				log.Printf("failed to parse: %s", err)
				continue
			}
			slashSplitDir := strings.Split(extractUrl.Path, "/")
			path := extractUrl.Hostname()
			for _, dir := range slashSplitDir[1:] {
				if strings.HasSuffix(dir, ".html") {
					break
				}
				path += "/" + dir
				if err := makeDirIfNotExist(path); err != nil {
					log.Printf("failed to make directory: %s", err)
				}
			}
			fmt.Printf("path: %s\n", path)
			if !strings.HasSuffix(path, ".html") {
				path += "/index.html"
			}
			if err := download(item, path); err != nil {
				log.Printf("failed to download %s to %s: %s", item, path, err)
				continue
			}
		}
	}
	fmt.Printf("depth %d %s\n", depth, urlStr)
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
				pURL, err := url.Parse(link)
				if err != nil {
					log.Printf("failed to parse %s: %s", link, err)
					continue
				}
				hostname := pURL.Hostname()
				if err := makeDirIfNotExist(hostname); err != nil {
					log.Printf("failed to make directory: %s", err)
				}
				path := hostname + "/index.html"
				if err := download(link, path); err != nil {
					log.Printf("failed to download %s to %s: %s", link, path, err)
					continue
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

func makeDirIfNotExist(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0644)
		if err != nil {
			log.Printf("failed to make %s: %v", dirName, err)
			return err
		}
		err = os.Chmod(dirName, 0744)
		if err != nil {
			log.Printf("failed to change permmision of %s: %v", dirName, err)
			return err
		}
	}
	return nil
}

func download(urlStr, path string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return fmt.Errorf("failed to Get %s: %s", urlStr, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to Get %s: %s", urlStr, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to read %s: %s", urlStr, err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s: %s", path, err)
	}
	return nil
}
