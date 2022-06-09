package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCustomSortBy1st2ndKey(t *testing.T) {
	var tests = []struct {
		key1, key2 string
		expected   []*Track
	}{
		{"Title", "Artist", []*Track{
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}},
		{"Title", "Album", []*Track{
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}},
		{"Title", "Year", []*Track{
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}},
		{"Length", "Year", []*Track{
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		}},
		{"Album", "Year", []*Track{
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}},
	}
	for _, test := range tests {
		actual := CustomSortBy1st2ndKey(test.key1, test.key2, tracks1)
		if !equal(test.expected, actual) {
			t.Errorf("Not expected.  expected %v, actual %v", test.expected, actual)
		}
	}
}

func equal(x, y []*Track) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i].Title != y[i].Title || x[i].Artist != y[i].Artist || x[i].Album != y[i].Album || x[i].Year != y[i].Year || x[i].Length != y[i].Length {
			return false
		}
	}
	return true
}

func TestCustomSortBy1st2ndKey1Time(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	CustomSortBy1st2ndKey(keys[0].key1, keys[0].key2, tracks100forCustom)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 Track was sorted by CustomSort. %d ns\n", nanoSec)
}

func TestCustomSortBy1st2ndKeyAlreadySorted(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	CustomSortBy1st2ndKey(keys[0].key1, keys[0].key2, tracks100SortedForCustom)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 sorted Track was sorted by CustomSort. %d ns\n", nanoSec)
}

func TestCustomSortBy1st2ndKeyReversed(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	CustomSortBy1st2ndKey(keys[0].key1, keys[0].key2, tracks100ReversedForCustom)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 reversed Track was sorted by CustomSort. %d ns\n", nanoSec)
}

func TestSortStable1Time(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	SortStable(keys[0].key1, keys[0].key2, tracks100forSortStable)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 Track was sorted by SortStable. %d ns\n", nanoSec)
}

func TestSortStableAlreadySorted(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	SortStable(keys[0].key1, keys[0].key2, tracks100SortedForSortStable)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 sorted Track was sorted by SortStable. %d ns\n", nanoSec)
}

func TestSortStableReversed(t *testing.T) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	start := time.Now()
	SortStable(keys[0].key1, keys[0].key2, tracks100ReversedForSortStable)
	nanoSec := time.Since(start).Nanoseconds()
	fmt.Printf("100 reversed Track was sorted by SortStable. %d ns\n", nanoSec)
}

func BenchmarkCustomSortBy1st2ndKey(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		CustomSortBy1st2ndKey(keys[i%length].key1, keys[i%length].key2, tracks100forCustom)
	}
}

func BenchmarkCustomSortBy1st2ndKeyAlreadySorted(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		CustomSortBy1st2ndKey(keys[i%length].key1, keys[i%length].key2, tracks100SortedForCustom)
	}
}

func BenchmarkSortStable(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		SortStable(keys[i%length].key1, keys[i%length].key2, tracks100forSortStable)
	}
}

func BenchmarkSortStableAlreadySorted(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		SortStable(keys[i%length].key1, keys[i%length].key2, tracks100SortedForSortStable)
	}
}

func BenchmarkCustomSortBy1st2ndKeyRandom(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
		{key1: "Album", key2: "Year"},
		{key1: "Length", key2: "Title"},
		{key1: "Artist", key2: "Album"},
		{key1: "Year", key2: "Length"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		CustomSortBy1st2ndKey(keys[i%length].key1, keys[i%length].key2, tracks100forCustom)
	}
}

func BenchmarkSortStableRandom(b *testing.B) {
	type Keys struct {
		key1 string
		key2 string
	}
	keys := []Keys{
		{key1: "Title", key2: "Artist"},
		{key1: "Album", key2: "Year"},
		{key1: "Length", key2: "Title"},
		{key1: "Artist", key2: "Album"},
		{key1: "Year", key2: "Length"},
	}
	length := len(keys)
	for i := 0; i < b.N; i++ {
		SortStable(keys[i%length].key1, keys[i%length].key2, tracks100forCustom)
	}
}
