// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	visit(nil, doc)
}

func visit(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		if n.Data == "img" || n.Data == "script" {
			for _, item := range n.Attr {
				if item.Key == "src" {
					fmt.Println(item.Val)
				}
			}
		}
		if n.Data == "link" {
			isStylesheetLink := false
			for _, item := range n.Attr {
				if item.Key == "rel" && item.Val == "stylesheet" {
					isStylesheetLink = true
				}
				if isStylesheetLink {
					if item.Key == "href" {
						fmt.Println(item.Val)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(stack, c)
	}
}

//!-
