// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

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

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item string)

	visitAll = func(item string) {
		fmt.Printf("visit %v\n", item)
		for _, v := range m[item] {
			visitAll(v)
			if !seen[v] {
				seen[v] = true
				order = append(order, v)
				fmt.Printf("append %v\n", v)
				break
			}
		}
		if !seen[item] {
			seen[item] = true
			order = append(order, item)
			fmt.Printf("append %v\n", item)

		}
	}
	for k := range m {
		visitAll(k)
	}
	return order
}
