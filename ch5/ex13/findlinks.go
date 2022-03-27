// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"ch5/ex13/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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
				pURL, err := url.Parse(item)
				if err != nil {
					log.Printf("failed to parse %s: %s", item, err)
					continue
				}
				hostname := pURL.Hostname()
				if err := makeDirIfNotExist(hostname); err != nil {
					log.Printf("failed to make directory: %s", err)
				}
				path := hostname + "/index.html"
				if err := download(item, path); err != nil {
					log.Printf("failed to download %s to %s: %s", item, path, err)
					continue
				}
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlStr string) []string {
	list, err := links.Extract(urlStr)
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
	return list
}

func makeDirIfNotExist(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0644)
		if err != nil {
			log.Printf("making %s: %v", dirName, err)
			return err
		}
		err = os.Chmod(dirName, 0744)
		if err != nil {
			log.Printf("changing permmision of %s: %v", dirName, err)
			return err
		}
	}
	return nil
}

func download(urlStr, path string) error {
	resp, err := http.Get(urlStr)
	if err != nil {
		return fmt.Errorf("gettig %s failed: %s", urlStr, err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s failed: %s", urlStr, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("reading %s failed: %s", urlStr, err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return fmt.Errorf("writing %s failed: %s", path, err)
	}
	return nil
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
