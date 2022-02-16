// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	length := len(os.Args)
	if length != 3 {
		fmt.Fprintf(os.Stderr, "invalid number of args")
		os.Exit(1)
	}
	outline(os.Args[1], os.Args[2])

}

func outline(url string, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement, id)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		node := ElementByID(n, id)
		if node != nil {
			return true
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		node := ElementByID(n, id)
		if node != nil {
			return true
		}
	}
	return false
}

//!-startend

func ElementByID(doc *html.Node, id string) *html.Node {
	if doc.Type == html.ElementNode {
		for _, script := range doc.Attr {
			if script.Key == "id" && script.Val == id {
				fmt.Printf("id MATCH %s\n", id)
				return doc
			}
		}
	}
	return nil
}
