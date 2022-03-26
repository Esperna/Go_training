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
	dispTextNode(nil, doc)
}

func dispTextNode(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
	}
	if !(n.Data == "style" || n.Data == "script") {
		if n.Type == html.TextNode {
			fmt.Printf("%s", n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dispTextNode(stack, c)
	}
}

//!-
