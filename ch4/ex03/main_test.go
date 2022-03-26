package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		array [6]int
		want  [6]int
	}{
		{[6]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
	}
	for _, test := range tests {
		reverse(&test.array)
		if test.array != test.want {
			t.Errorf("test.array != test.want. test.array is %v test.want is %v", test.array, test.want)
		}
	}
}
