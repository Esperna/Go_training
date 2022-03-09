package main

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		expected bool
		given    sort.Interface
	}{
		{true, IntSlice{1, 2, 1}},
		{false, IntSlice{1, 2, 3}},
		{true, StringSlice{"time", "after", "time"}},
		{false, StringSlice{"time", "goes", "by"}},
	}
	for _, test := range tests {
		actual := IsPalindrome(test.given)
		if actual != test.expected {
			t.Errorf("(%v): expected %t, actual %t", test.given, test.expected, actual)
		}
	}
}
