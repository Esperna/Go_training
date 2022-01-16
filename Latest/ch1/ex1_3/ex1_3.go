// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	Echo1()
	secs := time.Since(start).Seconds()
	fmt.Println(secs)

	start = time.Now()
	Echo2()
	secs = time.Since(start).Seconds()
	fmt.Println(secs)
}

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}