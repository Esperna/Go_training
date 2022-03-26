// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 96.

// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//!+
func main() {
	seen := wordfreq(os.Stdin)
	var names []string
	for w := range seen {
		names = append(names, w)
	}
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, seen[name])
	}
}

func wordfreq(r io.Reader) map[string]int {
	seen := make(map[string]int) // a set of strings
	input := bufio.NewScanner(r)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		seen[line]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	return seen
}

//!-
