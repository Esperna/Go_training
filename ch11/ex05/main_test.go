package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		want int
		s    string
		sep  string
	}{
		{3, "a:b:c", ":"},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got, want := len(words), test.want; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, want)
		}
	}
}
