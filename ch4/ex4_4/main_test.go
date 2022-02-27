package main

import "testing"

func TestRotate(t *testing.T) {
	var tests = []struct {
		inputs    []int
		rotNumber int
		wants     []int
	}{
		{[]int{1, 2, 3}, 1, []int{2, 3, 1}},
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
