// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"ch6/ex04/popcount"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) Len() int {
	sum := 0
	for _, word := range s.words {
		sum += popcount.PopCount(word)
	}
	return sum
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	for i, _ := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	ret.words = make([]uint64, len(s.words))
	copy(ret.words, s.words)
	return &ret
}

func (s *IntSet) AddAll(vals ...int) *IntSet {
	for _, v := range vals {
		s.Add(v)
	}
	return s
}

func (s *IntSet) IntersectWith(t *IntSet) {
	lengthS := len(s.words)
	lengthT := len(t.words)

	if lengthT < lengthS {
		for i, word := range t.words {
			s.words[i] &= word
		}
		for j := lengthT; j < lengthS; j++ {
			s.words[j] = 0
		}
	} else {
		for i, _ := range s.words {
			s.words[i] &= t.words[i]
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	lengthS := len(s.words)
	lengthT := len(t.words)

	if lengthT < lengthS {
		for i, word := range t.words {
			mask := s.words[i] & word
			s.words[i] &= ^mask
		}
	} else {
		for i, _ := range s.words {
			mask := t.words[i] & s.words[i]
			s.words[i] &= ^mask
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	lengthS := len(s.words)
	lengthT := len(t.words)

	if lengthT < lengthS {
		for i, word := range t.words {
			s.words[i] ^= word
		}
	} else {
		for i, _ := range s.words {
			s.words[i] ^= t.words[i]
		}
		for i := lengthS; i < lengthT; i++ {
			s.words = append(s.words, t.words[i])
		}
	}
}

func (s *IntSet) Elem() []uint64 {
	var elem []uint64
	var offset uint64
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			mask := uint64(1)
			if word&mask == 1 {
				elem = append(elem, offset+uint64(i))
			}
			word = word >> 1
		}
		offset += 64
	}
	return elem
}
