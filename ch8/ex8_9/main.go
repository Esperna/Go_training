// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type directoryInfo struct {
	size, fileNum int64
	index         int
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	numDirectry := len(roots)
	if numDirectry == 0 {
		roots = []string{"."}
		numDirectry = 1
	}
	fmt.Printf("numDirectry:%d\n", numDirectry)
	var dir = make([]string, numDirectry)

	//!+
	// Traverse each root of the file tree in parallel.
	dirInfoChs := make(chan directoryInfo)
	dirInfos := make([]directoryInfo, numDirectry)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		dir[i] = root
		go walkDir(root, &n, dirInfoChs, i)
	}
	go func() {
		n.Wait()
		close(dirInfoChs)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case dirInfoCh, ok := <-dirInfoChs:
			if !ok {
				break loop // dirInfoChs was closed
			}
			dirInfos[dirInfoCh.index].fileNum++
			dirInfos[dirInfoCh.index].size += dirInfoCh.size
		case <-tick:
			for i, dirInfo := range dirInfos {
				fmt.Println(i)
				printDiskUsageAtEachDir(dirInfo.fileNum, dirInfo.size, dir[dirInfo.index])
			}
		}
	}
	var nfiles, nbytes int64
	for _, dirInfo := range dirInfos {
		nfiles += dirInfo.fileNum
		nbytes += dirInfo.size
	}
	printDiskUsage(nfiles, nbytes) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func printDiskUsageAtEachDir(nfiles, nbytes int64, dir string) {
	fmt.Printf("%d files  %.1f GB %s\n", nfiles, float64(nbytes)/1e9, dir)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on dirInfoChs.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, dirInfoChs chan<- directoryInfo, index int) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, dirInfoChs, index)
		} else {
			dirInfoChs <- directoryInfo{entry.Size(), 1, index}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
