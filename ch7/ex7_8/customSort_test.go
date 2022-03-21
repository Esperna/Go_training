package main

import "testing"

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
		actual := customSortBy1st2ndKey(test.key1, test.key2)
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
