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
		actual, _ := sum(test.given...)
		if actual != test.expected {
			t.Errorf("(%v): expected %d, actual %d", test.given, test.expected, actual)
		}
	}
}

func TestInvalidArguments(t *testing.T) {
	var tests = []struct {
		expected int
		errStr   string
		given    []int
		f        func(...int) (int, error)
	}{
		{0, "no arguments", []int{}, sum},
		{0, "no arguments", []int{}, max},
		{0, "no arguments", []int{}, min},
	}
	for _, test := range tests {
		_, err := test.f(test.given...)
		if err == nil {
			t.Errorf("(%v): expected not nil, actual %v", test.given, err)
		} else {
			str := err.Error()
			if str != test.errStr {
				t.Errorf("(%v): expected %s, actual %s", test.given, test.errStr, str)
			}
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		expected int
		given    []int
	}{
		{5, []int{5}},
		{777, []int{1, 0, 777, 5}},
	}
	for _, test := range tests {
		actual, _ := max(test.given...)
		if actual != test.expected {
			t.Errorf("(%v): expected %d, actual %d", test.given, test.expected, actual)
		}

	}
}

func TestMin(t *testing.T) {
	var tests = []struct {
		expected int
		given    []int
	}{
		{5, []int{5}},
		{0, []int{1, 0, 777, 5}},
	}
	for _, test := range tests {
		actual, _ := min(test.given...)
		if actual != test.expected {
			t.Errorf("(%v): expected %d, actual %d", test.given, test.expected, actual)
		}
	}
}
