package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
		f        func(s string) string
	}{
		{"$foo", "3", Length},
		{"$foo$foo", "33", Length},
		{"$foo $foo", "3 3", Length},
		{"$hoge, $foo, $fuga", "4, 3, 4", Length},
		{"$hoge, $foo, $fuga", "HOGE, FOO, FUGA", strings.ToUpper},
	}

	for _, test := range tests {
		result := expand(test.input, test.f)
		if result != test.expected {
			t.Errorf("result is not expected. result is %s. expected is %s.", result, test.expected)
		}
	}
}
