package main

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		expected string
		given    string
	}{
		{"123", "123"},
		{"12,345", "12345"},
	}
	for _, test := range tests {
		actual := comma(test.given)
		if actual != test.expected {
			t.Errorf("(%s): expected %s, actual %s", test.given, test.expected, actual)
		}
	}

}
