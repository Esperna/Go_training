// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import (
	"fmt"
	"math"
)

//!+
func sum(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	total := 0
	for _, val := range vals {
		total += val
	}
	return total, nil
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	max := math.MinInt
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no arguments")
	}
	min := math.MaxInt
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min, nil
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
