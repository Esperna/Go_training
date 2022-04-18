package main

import (
	"fmt"
	"sync"
)

func main() {
	channelA := make(chan int)
	channelB := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := 0; x < 100; x++ {
			channelA <- x
			fmt.Printf("Ping%d\n", x)
			x = <-channelB
		}
		close(channelA)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range channelA {
			fmt.Printf("Pong%d\n", x)
			x++
			channelB <- x
		}
	}()
	wg.Wait()
}
