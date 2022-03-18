package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementById(t *testing.T) {
	var tests = []struct {
		node     html.Node
		id       string
		expected bool
	}{
		{html.Node{Parent: nil, FirstChild: nil, LastChild: nil, PrevSibling: nil, NextSibling: nil, Type: html.ElementNode, DataAtom: 0, Data: "", Namespace: "",
			Attr: []html.Attribute{{Namespace: "html", Key: "lang", Val: "en"}}}, "lang", true},
		{html.Node{Parent: nil, FirstChild: nil, LastChild: nil, PrevSibling: nil, NextSibling: nil, Type: html.ElementNode, DataAtom: 0, Data: "", Namespace: "",
			Attr: []html.Attribute{{Namespace: "html", Key: "lang", Val: "en"}}}, "id", false},
		{html.Node{Parent: nil, FirstChild: nil, LastChild: nil, PrevSibling: nil, NextSibling: nil, Type: html.DoctypeNode, DataAtom: 0, Data: "", Namespace: "",
			Attr: []html.Attribute{{Namespace: "", Key: "", Val: ""}}}, "lang", false},
	}
	for _, test := range tests {
		result := ElementByID(&test.node, test.id)
		if result != &test.node {
			t.Errorf("result is not expected. result is %v. expected is %v.", result, &test.node)
		}
	}
}

func TestElementByIdWithHtmlParse(t *testing.T) {
	resp, err := http.Get("https://golang.org")
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
	doc = ElementByID(doc, "lang")
}
