package main

import(
	"fmt"
	"time"
	"gopl.io/ch2/popcount_ex2_4"
)

func main(){
	const number uint64 = 1000
	fmt.Printf("InputNumber %d\n", number)
	displayFuncTime(popcount.PopCount, number)
	displayFuncTime(popcount.PopCountUsingFor, number)
	displayFuncTime(popcount.PopCountUsingInputBitShift, number)
}

type funcTemplate func(uint64) int

func displayFuncTime(f funcTemplate, t uint64) {
	start := time.Now()
	count := f(t)
	fmt.Printf("%d elapsed %dbit\n", time.Since(start), count)
}
//!-
