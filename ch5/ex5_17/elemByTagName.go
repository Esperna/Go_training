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
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http get failed: %v", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "html parse failed: %v", err)
			os.Exit(1)
		}
		images := ElementsByTagName(doc, "img")
		headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

		for _, n := range images {
			fmt.Printf("Data: %s\tNamespace: %s\tAttr: %v\n", n.Data, n.Namespace, n.Attr)
		}
		for _, n := range headings {
			fmt.Printf("Data: %s\tNamespace: %s\tAttr: %v\n", n.Data, n.Namespace, n.Attr)
		}
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodeList []*html.Node
	nodeList = forEachNode(doc, startElement, endElement, name, nodeList)
	return nodeList
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node, names []string) bool, names []string, nodeList []*html.Node) []*html.Node {
	if pre != nil {
		if pre(n, names) {
			nodeList = append(nodeList, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodeList = forEachNode(c, pre, post, names, nodeList)
	}

	if post != nil {
		post(n, names)
	}
	return nodeList
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, names []string) bool {
	if n.Type == html.ElementNode {
		//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		for _, name := range names {
			if name == n.Data {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, names []string) bool {
	if n.Type == html.ElementNode {
		depth--
		//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return false
}

//!-startend
