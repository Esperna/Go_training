//Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("%fsec",secs)
}
