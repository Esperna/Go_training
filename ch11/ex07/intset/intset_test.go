// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	if x.String() != "{1}" {
		t.Errorf("x.String() != {1}. x.String() is %s", x.String())
	}
	x.Add(144)
	if x.String() != "{1 144}" {
		t.Errorf("x.String() != {1 144}. x.String() is %s", x.String())
	}
	x.Add(9)
	if x.String() != "{1 9 144}" {
		t.Errorf("x.String() != {1 9 144}. x.String() is %s", x.String())
	}
}

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	if !x.Has(9) {
		t.Errorf("x.Has(9) != true. x.Has(9) is %t", x.Has(9))
	}
	if x.Has(123) {
		t.Errorf("x.Has(123) != false. x.Has(123) is %t", x.Has(123))

	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	if x.String() != "{1 9 42 144}" {
		t.Errorf("x.String() != {1 9 42 144}. x.String() is %s", x.String())
	}

	var a, b IntSet
	a.Add(1)
	a.Add(144)
	a.Add(9)
	b.Add(9)
	b.Add(42)
	b.Add(100)
	b.Add(1000)
	a.UnionWith(&b)
	if a.String() != "{1 9 42 100 144 1000}" {
		t.Errorf("a.String() != {1 9 42 100 144 1000}. a.String() is %s", a.String())
	}
}

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if x.Len() != 4 {
		t.Errorf("x.Len() != 4, x.Len() is %d", x.Len())
	}
	x.Add(256)
	if x.Len() != 5 {
		t.Errorf("x.Len() != 5, x.Len() is %d", x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.Remove(9)
	if x.String() != "{1 42 144}" {
		t.Errorf("x.String() != {1 42 144}, x.String() is %s", x.String())
	}
	x.Remove(9999)
	if x.String() != "{1 42 144}" {
		t.Errorf("x.String() != {1 42 144}, x.String() is %s", x.String())
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.Clear()
	if x.String() != "{}" {
		t.Errorf("x.String() != {}, x.String() is %s", x.String())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	var y *IntSet
	y = x.Copy()
	if y.String() != x.String() {
		t.Errorf("y.String() != x.String(), y.String() is %s x.String() is %s", y.String(), x.String())
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 2, 3)
	if x.String() != "{1 2 3}" {
		t.Errorf("x.String() != {1 2 3}, x.String() is %s", x.String())
	}
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)

	x.IntersectWith(&y)
	if x.String() != "{9}" {
		t.Errorf("x.String() != {9}. x.String() is %s", x.String())
	}

	var a, b IntSet
	a.Add(1)
	a.Add(144)
	a.Add(9)
	b.Add(9)
	b.Add(42)
	b.Add(144)
	b.Add(256)
	a.IntersectWith(&b)
	if a.String() != "{9 144}" {
		t.Errorf("a.String() != {9 144}. a.String() is %s", a.String())
	}

}
func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)

	x.DifferenceWith(&y)
	if x.String() != "{1 144}" {
		t.Errorf("x.String() != {1 144}. x.String() is %s", x.String())
	}

	var a, b IntSet
	a.Add(1)
	a.Add(144)
	a.Add(9)
	b.Add(9)
	b.Add(42)
	b.Add(144)
	b.Add(256)
	a.DifferenceWith(&b)
	if a.String() != "{1}" {
		t.Errorf("a.String() != {1}. a.String() is %s", a.String())
	}
}
func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)

	x.SymmetricDifference(&y)
	if x.String() != "{1 42 144}" {
		t.Errorf("x.String() != {1 42 144}. x.String() is %s", x.String())
	}

	var a, b IntSet
	a.Add(1)
	a.Add(144)
	a.Add(9)
	b.Add(9)
	b.Add(42)
	b.Add(144)
	b.Add(256)
	a.SymmetricDifference(&b)
	if a.String() != "{1 42 256}" {
		t.Errorf("a.String() != {1 42 256}. a.String() is %s", a.String())
	}

}

var rng *rand.Rand

func init() {
	seed := time.Now().UTC().UnixNano()
	log.Printf("Random seed: %d", seed)
	rng = rand.New(rand.NewSource(seed))
}

func BenchmarkAdd(b *testing.B) {
	var x IntSet
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			x.Add(rng.Intn(0x1000))
		}
	}
}
