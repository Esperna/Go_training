package main

import "testing"

func TestSum(t *testing.T) {
	var tests = []struct {
		expected int
		given    []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{1, []int{0, 1}},
		{10, []int{1, 2, 3, 4}},
	}
	for _, test := range tests {
		actual := sum(test.given...)
		if actual != test.expected {
			t.Errorf("(%v): expected %d, actual %d", test.given, test.expected, actual)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		expected int
		given    []int
	}{
		{5, []int{5}},
		{777, []int{1, 0, 777, 0}},
	}
	for _, test := range tests {
		actual := max(test.given...)
		if actual != test.expected {
			t.Errorf("(%v): expected %d, actual %d", test.given, test.expected, actual)
		}

	}
}
