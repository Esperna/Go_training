package main

import "testing"

func TestToposortByMap(t *testing.T) {
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
