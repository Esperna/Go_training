package main

import (
	"os"
	"testing"
)

func TestCountUnicode(t *testing.T) {
	var tests = []struct {
		expected unicodeCount
		given    string
	}{
		{unicodeCount{28, 10, 29}, "in"},
	}
	for _, test := range tests {
		f, err := os.Open(test.given)
		if err != nil {
			t.Errorf("%s", err)
		}
		actual := countUnicode(f)
		if actual != test.expected {
			t.Errorf("input file name %s: expected %v, actual %v", test.given, test.expected, actual)
		}
	}

}
