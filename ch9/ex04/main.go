package main

import (
	"flag"
	"fmt"
)

var n = flag.Int("n", 1, "number of go routine")

func main() {
	flag.Parse()
	fmt.Printf("number of goroutine is %d\n", *n)

	channels := make([]chan int, *n)
	for i := 0; i < *n; i++ {
		channels[i] = make(chan int)
	}

	go func() {
		for x := 0; ; x++ {
			channels[0] <- x
		}
	}()
	for i := 0; i < *n-1; i++ {
		go func(i int) {
			for {
				x := <-channels[i]
				channels[i+1] <- x + 1
			}
		}(i)
	}

	for x := range channels[*n-1] {
		fmt.Println(x)
	}
}
