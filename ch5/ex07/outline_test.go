package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	var tests = []struct {
		url  string
		want error
	}{
		{"https://golang.org/", nil},
		{"http://www.oceanus.in/", nil},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("outline(%s)", test.url)
		out = new(bytes.Buffer)
		if err := outline(test.url); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		_, got := html.Parse(strings.NewReader(out.(*bytes.Buffer).String()))
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}

	}

}
