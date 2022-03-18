// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := flag.String("url", "", "URL")
	id := flag.String("id", "", "html attribute key")
	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html GET failed: %s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse failed: %s", err)
		os.Exit(1)
	}
	ElementByID(doc, *id)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, id string) bool, id string) bool {
	if pre != nil {
		if pre(n, id) {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post, id) {
			return false
		}
	}
	if post != nil {
		if post(n, id) {
			return false
		}
	}
	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	findAttributeId := func(n *html.Node, id string) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == id {
					return true
				}
			}
		}
		return false
	}
	forEachNode(doc, findAttributeId, nil, id)
	return doc
}
