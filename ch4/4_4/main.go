// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import "fmt"

func main() {
	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	//reverse(s[:2])
	//reverse(s[2:])
	//reverse(s)
	l := len(s)
	k := 2 //rotate number
	s = append(s[k:l], s[0:k]...)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
