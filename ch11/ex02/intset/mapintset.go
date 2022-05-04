// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type Uint64Slice []uint64

func (s Uint64Slice) Len() int           { return len(s) }
func (s Uint64Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type MapIntSet struct {
	words map[uint64]bool
}

func (s *MapIntSet) Has(x int) bool {
	return s.words[uint64(x)]
}

func (s *MapIntSet) Add(x int) {
	s.words[uint64(x)] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	for key := range t.words {
		s.words[key] = true
	}
}

func (s *MapIntSet) String() string {
	var values Uint64Slice
	for value := range s.words {
		values = append(values, value)
	}
	sort.Sort(values)
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, value := range values {
		fmt.Fprintf(&buf, "%d", value)
		if i < len(values)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *MapIntSet) Len() int {
	return len(s.words)
}

func (s *MapIntSet) Remove(x int) {
	delete(s.words, uint64(x))
}
