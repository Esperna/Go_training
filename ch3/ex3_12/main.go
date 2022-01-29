package main

import (
	"fmt"
	"os"
	"sort"
)

type RuneSlice []rune

func (r RuneSlice) Len() int { return len(r) }

func (r RuneSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r RuneSlice) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r RuneSlice = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func main() {
	if len(os.Args) == 3 {
		if SortStringByCharacter(os.Args[1]) == SortStringByCharacter(os.Args[2]) {
			fmt.Println("Anagram!")
		} else {
			fmt.Println("Not Anagram!")
		}
	} else {
		fmt.Println("Invalid Number of Argument")
	}
}