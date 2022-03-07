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
func sum(vals ...int) (total int, err error) {
	total = 0
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	for _, val := range vals {
		total += val
	}
	return total, err
}

func max(vals ...int) (max int, err error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	max = math.MinInt
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max, err
}

func min(vals ...int) int {
	length := len(vals)
	if length == 0 {
		fmt.Fprintf(os.Stderr, "no vals:")
		return 0
	}
	var min int = math.MaxInt
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

//!-

func main() {
	fmt.Println("sum()")
	fmt.Println(sum()) //  "0"
	fmt.Println("sum(3)")
	fmt.Println(sum(3)) //  "3"
	fmt.Println("sum(1, 2, 3, 4)")
	fmt.Println(sum(1, 2, 3, 4)) //  "10"

	fmt.Println("values := []int{1, 2, 3, 4}")
	values := []int{1, 2, 3, 4}
	fmt.Println("sum(values...)")
	fmt.Println(sum(values...)) // "10"

	fmt.Println("max()")
	fmt.Println(max())
	fmt.Println("max(5)")
	fmt.Println(max(5))
	fmt.Println("max(1,0,777,10)")
	fmt.Println(max(1, 0, 777, 10))

	fmt.Println("values := []int{1, 2, 3, 4}")
	fmt.Println("max(values...)")
	fmt.Println(max(values...)) // "4"

	fmt.Println("min()")
	fmt.Println(min())
	fmt.Println("min(5)")
	fmt.Println(min(5))
	fmt.Println("min(1,0,777,10)")
	fmt.Println(min(1, 0, 777, 10))

	fmt.Println("values := []int{1, 2, 3, 4}")
	fmt.Println("min(values...)")
	fmt.Println(min(values...)) // "1
}
