package main

import "testing"

func TestToposortBySlice(t *testing.T) {
	// var tests = []struct {
	// 	input1, input2 []byte
	// 	want           int
	// }{
	// 	{[]byte("x"), []byte("x"), 0},
	// 	{[]byte("X"), []byte("X"), 0},
	// 	{[]byte("HelloWorld"), []byte("HelloWorld"), 0},
	// }
	// for _, test := range tests {
	// 	sum1 := sha256.Sum256(test.input1)
	//
	// }
	isExpected := true
	courses := topoSort(prereqs)
	for i := 0; i < len(courses)-1; i++ {
		for _, prereq := range prereqs[courses[i]] {
			if courses[i+1] == prereq {
				isExpected = false
				t.Errorf("next course is prereq. current: %s, next: %s", courses[i], courses[i+1])
			}
		}
	}
	if !isExpected {
		t.Errorf("Not Expected toposort result.")
	}
}
