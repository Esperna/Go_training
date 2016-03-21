// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"fmt"
	"os"
)

//!+
import "crypto/sha256"
import "crypto/sha512"

func main() {
	length := len(os.Args)
	if length > 2 {
		if os.Args[2] == "-384" {
			c1 := sha512.Sum384([]byte(os.Args[1]))
			fmt.Printf("%s\n%x\n%T\n", os.Args[1], c1, c1)
		} else if os.Args[2] == "-512" {
			c1 := sha512.Sum512([]byte(os.Args[1]))
			fmt.Printf("%s\n%x\n%T\n", os.Args[1], c1, c1)
		} else {
			c1 := sha256.Sum256([]byte(os.Args[1]))
			fmt.Printf("%s\n%x\n%T\n", os.Args[1], c1, c1)
		}
	} else if length == 2 {
		c1 := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("%s\n%x\n%T\n", os.Args[1], c1, c1)
	} else {
		fmt.Printf("less arguments")
	}
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
