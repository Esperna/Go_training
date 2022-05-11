// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package cyclic

import (
	"testing"
)

func TestIsCyclicLinkedList(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	d, e := &link{value: "d"}, &link{value: "e"}
	d.tail, e.tail = e, nil

	var tests = []struct {
		given *link
		want  bool
	}{
		{a, true},
		{b, true},
		{c, true},
		{d, false},
	}
	for _, test := range tests {
		actual := IsCyclic(test.given)
		if actual != test.want {
			t.Errorf("actual:%t want:%t", actual, test.want)
		}
	}
}

func TestIsCyclicSlice(t *testing.T) {
	type Slice struct {
		slices []Slice
	}
	s := make([]Slice, 1)
	s[0].slices = s
	actual := IsCyclic(s)
	want := true
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}
}

func TestIsCyclicArray(t *testing.T) {
	type pointer struct {
		next [1]*pointer
	}

	a1, a2, a3 := pointer{}, pointer{}, pointer{}
	a1.next, a2.next, a3.next = [1]*pointer{&a2}, [1]*pointer{&a3}, [1]*pointer{&a1}
	actual := IsCyclic(a1)
	want := true
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}
}

/*
func TestIsCyclicMap(t *testing.T) {
	type pointer struct {
		p map[string]pointer
	}

	p1 := pointer{make(map[string]pointer)}
	p1.p["p"] = p1
	actual := IsCyclic(p1)
	want := true
	if actual != want {
		t.Errorf("actual:%t want:%t", actual, want)
	}

}
*/
