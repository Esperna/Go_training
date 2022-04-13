package main

import (
	"flag"
	"fmt"
	"time"
)

var n = flag.Int("n", 1, "number of go routine")
var m = flag.Int("m", 1, "number of sending trigger")

func main() {
	flag.Parse()
	fmt.Printf("number of goroutine is %d\n", *n)
	fmt.Printf("number of sending trigger is %d\n", *m)

	channels := make([]chan int, *n)
	for i := 0; i < *n; i++ {
		channels[i] = make(chan int)
	}
	done := make(chan struct{})

	var t time.Time
	go func() {
		t = time.Now()
		for x := 0; x < *m; x++ {
			channels[0] <- x
		}
		<-done
	}()
	for i := 0; i < *n-1; i++ {
		go func(i int) {
			for j := 0; j < *m; j++ {
				x := <-channels[i]
				channels[i+1] <- x + 1
			}
			<-done
		}(i)
	}

	for i := 0; i < *m; i++ {
		x := <-channels[*n-1]
		fmt.Printf("%dμs\n", time.Since(t).Microseconds())
		fmt.Println(x)
	}
	close(done)
}
