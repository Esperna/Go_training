package main

import "testing"

func TestJoin(t *testing.T) {
	var tests = []struct {
		expected string
		given1   string
		given2   []string
	}{
		{"ab cd ef", " ", []string{"ab", "cd", "ef"}},
		{"abcdef", "", []string{"ab", "cd", "ef"}},
	}
	for _, test := range tests {
		actual := Join(test.given1, test.given2...)
		if actual != test.expected {
			t.Errorf("(%s %s): expected %s, actual %s", test.given1, test.given2, test.expected, actual)
		}
	}
}
