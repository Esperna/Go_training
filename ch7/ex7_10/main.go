package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	return true
}
