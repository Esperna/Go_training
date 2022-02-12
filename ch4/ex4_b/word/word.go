// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 305.
//!+

// Package word provides utilities for word games.
package word

import (
	"fmt"
)

var stringArraysA [38]string = [38]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December", "a", "ab", "abc", "abcd", "abcde", "abcdef", "abcdefg", "abcdefgh", "abcdefghi", "abcdefghij", "abcdefghijk", "abcdefghijkl", "abcdefghijklm", "abcdefghijklmn", "o", "op", "opqr", "s", "t", "u", "vvv", "wwww", "xxxxxx", "xyy", "xyz"}
var stringArraysB [38]string = [38]string{"February", "January", "April", "June", "July", "August", "March", "September", "November", "October", "December", "May", "abcdef", "abcdefg", "opqr", "s", "t", "u", "vvv", "wwww", "xxxxxx", "abcdefghij", "abcdefghijk", "abcdefghijkl", "a", "ab", "abc", "abcd", "abcde", "abcdefghijklm", "abcdefghijklmn", "xyy", "abcdefgh", "abcdefghi", "o", "op", "xyz"}

func PrintSliceMacthUnefficient() {
	sliceA := stringArraysA[:]
	sliceB := stringArraysB[:]

	for _, s := range sliceB {
		for _, q := range sliceA {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}

func PrintSliceMacthEfficient() {
	sliceA := stringArraysA[:]
	sliceB := stringArraysB[:]
	alreadyAppeared := make(map[string]bool)

	for _, v := range sliceB {
		alreadyAppeared[v] = true
	}
	for _, v := range sliceA {
		if alreadyAppeared[v] {
			fmt.Printf("%s appears in both\n", v)
		}
	}
}

//!-
