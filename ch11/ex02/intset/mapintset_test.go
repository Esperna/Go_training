// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"testing"
)

func TestAddMapIntSet(t *testing.T) {
	var x1 IntSet
	var x2 MapIntSet

	x1.Add(1)
	x1.Add(144)
	x1.Add(9)

	x2.words = make(map[uint64]bool)
	x2.Add(1)
	x2.Add(144)
	x2.Add(9)

	if x1.String() != x2.String() {
		t.Errorf("x1:%s, x2:%s", x1.String(), x2.String())
	}
}

func TestUnionWithMapIntSet(t *testing.T) {
	var x1, y1 IntSet
	var x2, y2 MapIntSet

	x1.Add(1)
	x1.Add(144)
	x1.Add(9)
	y1.Add(9)
	y1.Add(42)
	x1.UnionWith(&y1)

	x2.words = make(map[uint64]bool)
	y2.words = make(map[uint64]bool)
	x2.Add(1)
	x2.Add(144)
	x2.Add(9)
	y2.Add(9)
	y2.Add(42)
	x2.UnionWith(&y2)

	if x1.String() != x2.String() {
		t.Errorf("Want:x1==x2 Actual: x1:%s, x2:%s", x1.String(), x2.String())
	}
}

func TestHasMapIntSet(t *testing.T) {
	var x1 IntSet
	var x2 MapIntSet

	x1.Add(1)
	x1.Add(144)
	x1.Add(9)

	x2.words = make(map[uint64]bool)
	x2.Add(1)
	x2.Add(144)
	x2.Add(9)

	for key := range x2.words {
		if x1.Has(int(key)) != x2.Has(int(key)) {
			t.Errorf("Want:x1.Has(key)==x2.Has(key) Actual: x1.Has(key)=%t, x2.Has(key)=%t, key=%d", x1.Has(int(key)), x2.Has(int(key)), key)
		}
	}
}
