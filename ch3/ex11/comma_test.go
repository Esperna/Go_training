package main

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		expected string
		given    string
	}{
		{"123", "123"},
		{"12,345", "12345"},
		{"123,456.789", "123456.789"},
		{"+1,234,567.89", "+1234567.89"},
		{"-123.456789", "-123.456789"},
		{".12", ".12"},
		{".1234", ".1234"},
		{"", ""},
	}
	for _, test := range tests {
		actual := comma(test.given)
		if actual != test.expected {
			t.Errorf("(%s): expected %s, actual %s", test.given, test.expected, actual)
		}
	}

}
