package main

import(
	"fmt"
	"time"
	"gopl.io/ch2/popcount_ex2_3"
)

func main(){
	var number uint64 = 1000
	fmt.Printf("InputNumber  %d\n", number)
	count := 0
	start := time.Now()
	count = popcount.PopCount(number)
	fmt.Printf("%d elapsed  %dbit\n", time.Since(start), count)
	count = 0
	start = time.Now()
	count = popcount.PopCountUsingFor(number)
	fmt.Printf("%d elapsed  %dbit\n", time.Since(start), count)
}

//!-
