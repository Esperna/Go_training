package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := [][]int{
		{80, 70, 95},
		{50, 90, 70},
		{80, 90, 70},
		{100, 50, 65},
	}
	fmt.Println(scores)
	sort.Sort(TestScore(scores))
	fmt.Println(scores)
}

type TestScore [][]int

func (s TestScore) Len() int {
	return len(s)
}

var FirstKey = 0
var SecondKey = 1

func (s TestScore) Less(i, j int) bool {
	if s[i][FirstKey] < s[j][FirstKey] {
		return true
	} else if s[i][FirstKey] == s[j][FirstKey] {
		return s[i][SecondKey] < s[j][SecondKey]
	} else {
		return false
	}
}

func (s TestScore) Swap(i, j int) {
	s[i][FirstKey], s[j][FirstKey] = s[j][FirstKey], s[i][FirstKey]
}
