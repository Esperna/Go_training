// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var duplines []string
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			duplines = append(duplines, line)
		}
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		dispFilenamesIncludingDuplicatedLines(f, duplines, arg)
		f.Close()
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dispFilenamesIncludingDuplicatedLines(f *os.File, duplines []string, filename string) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		for _, line := range duplines {
			if input.Text() == line && counts[line] == 0 {
				fmt.Printf("%s\t%s\n", line, filename)
				counts[line]++
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
