// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
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
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s", depth*2, "", n.Data)
	} else if n.Type == html.ElementNode {
		str := n.Data
		for _, a := range n.Attr {
			str = fmt.Sprintf("%s %s=\"%s\"", str, a.Key, a.Val)
		}
		if n.FirstChild == nil && (n.Data == "img" || n.Data == "meta" || n.Data == "link" || n.Data == "path") {
			fmt.Printf("%*s<%s", depth*2, "", str)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", str)
		}
		depth++
	} else if n.Type == html.TextNode && !isAllSpace(n.Data) {
		fmt.Printf("%*s%s\n", depth*2, "", strings.TrimSpace(n.Data))
	}
}

func endElement(n *html.Node) {
	if n.Type == html.CommentNode {
		fmt.Printf("-->\n")
	} else if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil && (n.Data == "img" || n.Data == "meta" || n.Data == "link" || n.Data == "path") {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend

func isAllSpace(s string) bool {
	r := []rune(s)
	for _, v := range r {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}
