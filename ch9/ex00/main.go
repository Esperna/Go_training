// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 218.

// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"time"
)

//!+
func main() {
	go spinner(100 * time.Millisecond)
	const n = 145
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		var s string
		for i := 0; i < 10; i++ {
			for _, r := range `-\|/ ` {
				fmt.Printf("\r%s%c", s, r)
				time.Sleep(delay)
			}
			s += " "
		}
		for i := 0; i < 10; i++ {
			for _, r := range `-\|/ ` {
				fmt.Printf("\r%s%c", s, r)
				time.Sleep(delay)
			}
			slice := []rune(s)
			s = string(slice[:(len(slice) - 1)])
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//!-
