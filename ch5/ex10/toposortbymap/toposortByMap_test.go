package main

import (
	"fmt"
	"testing"
)

func TestToposortByMap(t *testing.T) {
	isExpected := true
	courses := topoSort(prereqs)
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	for u := 0; u < len(courses); u++ {
		for v := u + 1; v < len(courses); v++ {
			isprereq := isPrereq(courses[u], courses[v])
			if isprereq && (u >= v) {
				isExpected = false
				t.Errorf("%s is not prereq of %s", courses[u], courses[v])
			}
		}
	}
	if !isExpected {
		t.Errorf("Not Expected toposort result.")
	}
}

func isPrereq(course1, course2 string) bool {
	if course1 == course2 {
		return false
	}
	for course := range prereqs[course2] {
		if course1 == course {
			return true
		}
		if isPrereq(course1, course) {
			return true
		}
	}
	return false

}
