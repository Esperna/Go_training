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
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
				}
			}
		}
		if n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					fmt.Println(img.Val)
				}
			}
		}
		if n.Data == "script" {
			for _, script := range n.Attr {
				if script.Key == "src" {
					fmt.Println(script.Val)
				}
			}
		}
		if n.Data == "link" {
			isStylesheetLink := false
			for _, link := range n.Attr {
				if link.Key == "rel" && link.Val == "stylesheet" {
					isStylesheetLink = true
				}
				if isStylesheetLink {
					if link.Key == "href" {
						fmt.Println(link.Val)
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
