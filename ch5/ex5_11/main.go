// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	counts := make(map[string]int)
	var visitAll func(items []string) bool

	visitAll = func(items []string) bool {
		for _, item := range items {
			fmt.Fprintf(os.Stdout, "%s\n", item)
			counts[item]++
			if counts[item] > 1 {
				fmt.Fprintf(os.Stderr, "graph is cyclic: %s appears again\n", item)
				return false
			}
			if !seen[item] {
				seen[item] = true
				if !visitAll(m[item]) {
					return false
				}
				order = append(order, item)
			}
			for k, _ := range counts {
				counts[k] = 0
			}
		}
		return true
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if !visitAll(keys) {
		return nil
	}
	return order
}

//!-main
