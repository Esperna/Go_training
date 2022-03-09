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
		{true, customSort{
			[]Track{
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			},
		},
		},
		{false, customSort{
			[]Track{
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			},
		},
		},
	}
	for _, test := range tests {
		actual := IsPalindrome(test.given)
		if actual != test.expected {
			t.Errorf("(%v): expected %t, actual %t", test.given, test.expected, actual)
		}
	}
}
