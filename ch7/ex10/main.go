package main

import (
	"sort"
	"time"
)

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type customSort struct {
	t []Track
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool {
	if x.t[i].Title != x.t[j].Title {
		return x.t[i].Title < x.t[j].Title
	}
	if x.t[i].Year != x.t[j].Year {
		return x.t[i].Year < x.t[j].Year
	}
	if x.t[i].Length != x.t[j].Length {
		return x.t[i].Length < x.t[j].Length
	}
	return false
}
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}

func main() {

}
