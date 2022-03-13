package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestStartElement(t *testing.T) {
	var tests = []struct {
		node     html.Node
		id       string
		expected bool
	}{
		{html.Node{nil, nil, nil, nil, nil, html.ElementNode, 0, "", "", []html.Attribute{html.Attribute{"html", "lang", "en"}}}, "lang", true},
		{html.Node{nil, nil, nil, nil, nil, html.ElementNode, 0, "", "", []html.Attribute{html.Attribute{"html", "lang", "en"}}}, "id", false},
		{html.Node{nil, nil, nil, nil, nil, html.DoctypeNode, 0, "", "", []html.Attribute{html.Attribute{"", "", ""}}}, "lang", false},
	}
	for _, test := range tests {
		result := startElement(&test.node, test.id)
		if result != test.expected {
			t.Errorf("result is not expected. result is %t. expected is %t.", result, test.expected)
		}
	}
}
