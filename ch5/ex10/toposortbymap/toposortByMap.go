// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms": {
		"data structures": true,
	},
	"calculus": {
		"linear algebra": true,
	},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},
	"data structures": {
		"discrete math": true,
	},
	"databases": {
		"data structures": true,
	},
	"discrete math": {
		"intro to programming": true,
	},
	"formal languages": {
		"discrete math": true,
	},
	"networks": {
		"operating systems": true,
	},
	"operating systems": {
		"data structures":       true,
		"computer organization": true,
	},
	"programming languages": {
		"data structures":       true,
		"computer organization": true,
	},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(m1 map[string]bool)

	visitAll = func(maps map[string]bool) {
		for k := range maps {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				order = append(order, k)
			}
		}
	}

	maps := make(map[string]bool)
	for key, values := range m {
		for _, v := range values {
			maps[key] = v
		}
	}
	visitAll(maps)
	return order
}
