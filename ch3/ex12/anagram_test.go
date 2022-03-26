package main

import "testing"

func TestSortStringByCharacter(t *testing.T) {
	var tests = []struct {
		expected       bool
		given1, given2 string
	}{
		{false, "ab", "ab"},
		{true, "abcde", "cbaed"},
		{false, "ab", "bac"},
		{true, "あいう", "ういあ"},
		{true, "あiう", "うiあ"},
	}
	for _, test := range tests {
		actual := isAnagram(test.given1, test.given2)
		if actual != test.expected {
			t.Errorf("(%s %s): expected %t, actual %t", test.given1, test.given2, test.expected, actual)
		}
	}
}
