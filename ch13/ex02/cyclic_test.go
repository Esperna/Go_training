// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package cyclic

import (
	"testing"
)

func TestIsCyclic(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a := &link{value: "a"}
	a.tail = a

	actual := IsCyclic(a)
	want := true
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}
}

func TestIsCyclic2(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	actual := IsCyclic(a)
	want := true
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}
}

func TestIsCyclic3(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, c, nil

	actual := IsCyclic(a)
	want := false
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}
}
