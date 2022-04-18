package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channelA := make(chan int)
	channelB := make(chan int)
	var wg sync.WaitGroup

	var x int
	wg.Add(1)
	go func() {
		defer wg.Done()
		tick := time.Tick(1 * time.Second)
	loop:
		for {
			select {
			case <-tick:
				break loop
			default:
				channelA <- x
				fmt.Printf("Ping\n")
				x = <-channelB
			}
		}
		close(channelA)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range channelA {
			fmt.Printf("Pong\n")
			x++
			channelB <- x
		}
	}()
	wg.Wait()
	fmt.Printf("total:%d\n", x)
}
