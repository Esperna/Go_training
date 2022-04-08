package main

import (
	"os"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		expected charCount
		given    string
	}{
		{charCount{3, 3, nil}, "in1"},
	}
	for _, test := range tests {
		f, err := os.Open(test.given)
		if err != nil {
			t.Errorf("%s", err)
		}
		actual := countChar(f)
		if actual != test.expected {
			t.Errorf("(%s): expected %v, actual %v", test.given, test.expected, actual)
		}
	}

}
