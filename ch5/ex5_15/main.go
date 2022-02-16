// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"math"
	"os"
)

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	length := len(vals)
	if length == 0 {
		fmt.Fprintf(os.Stderr, "no vals:")
		return 0
	}
	var max int = math.MinInt
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

//!-

func main() {
	//!+main
	fmt.Println("sum()")
	fmt.Println(sum()) //  "0"
	fmt.Println("sum(3)")
	fmt.Println(sum(3)) //  "3"
	fmt.Println("sum(1, 2, 3, 4)")
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	fmt.Println("values := []int{1, 2, 3, 4}")
	values := []int{1, 2, 3, 4}
	fmt.Println("sum(values...)")
	fmt.Println(sum(values...)) // "10"
	//!-slice

	fmt.Println("max()")
	fmt.Println(max())
	fmt.Println("max(5)")
	fmt.Println(max(5))
	fmt.Println("max(1,0,777,10)")
	fmt.Println(max(1, 0, 777, 10))

	fmt.Println("values := []int{1, 2, 3, 4}")
	fmt.Println("max(values...)")
	fmt.Println(max(values...)) // "10"
}
