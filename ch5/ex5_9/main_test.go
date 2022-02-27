package main

import "testing"

func TestExpand(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"$foo", "3"},
		{"$foo$foo", "33"},
		{"$foo $foo", "3 3"},
	}

	for _, test := range tests {
		result := expand(test.input, Length)
		if result != test.expected {
			t.Errorf("result is not expected. result is %s. expected is %s.", result, test.expected)
		}
	}
}
