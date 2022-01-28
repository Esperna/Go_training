// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	array := [...]string{"abc", "def", "hij", "klm", "nop", "qrs", "tuv", "wxy", "zab", "cde", "fgh", "ijk", "lmn", "opq", "rst", "uvw", "xyz"}
	start := time.Now()
	Echo1(array[:])
	secs := time.Since(start).Seconds()
	fmt.Println(secs)

	start = time.Now()
	Echo2(array[:])
	secs = time.Since(start).Seconds()
	fmt.Println(secs)

	start = time.Now()
	Echo2(array[:])
	secs = time.Since(start).Seconds()
	fmt.Println(secs)
}

func Echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func Echo3(args []string) {
	var s, sep string
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
