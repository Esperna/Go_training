package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

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

func crawl(path string) []string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	fmt.Println(path)
	var list []string
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadDir failed: %s", err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			dirpath := path + dir.Name() + "/"
			list = append(list, dirpath)
			fmt.Printf("%s\n", dirpath)
		} else {
			fmt.Printf("%s%s\n", path, dir.Name())
		}
	}
	return list
}
