package main

import "testing"

func TestLeftRotate(t *testing.T) {
	var tests = []struct {
		inputs    []int
		rotNumber int
		wants     []int
	}{
		{[]int{1, 2, 3}, 1, []int{2, 3, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5, []int{6, 7, 8, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -1, []int{3, 1, 2}},
	}
	for _, test := range tests {
		outputs := rotate(test.inputs, test.rotNumber)
		isMatch := true
		for i, _ := range outputs {
			if outputs[i] != test.wants[i] {
				isMatch = false
			}
		}
		if !isMatch {
			t.Errorf("rotate(inputs, rotNumber)!= test.want. outputs are %v wants are %v", outputs, test.wants)
		}
	}

}
