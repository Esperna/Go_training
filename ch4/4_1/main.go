// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"fmt"
	"gopl.io/ch2/popcount"
)

//!+
import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var temp uint8 = 0
	count := 0

	fmt.Println(len(c1))
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			temp = ^(uint8(c1[i]) | uint8(c2[i]))
			fmt.Println(temp)
			count += popcount.PopCount(uint64(temp))
		}
	}
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Printf("%d\n", count)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
