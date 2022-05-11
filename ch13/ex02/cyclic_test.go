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
