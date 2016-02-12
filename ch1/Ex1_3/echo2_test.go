package main

import (
	"testing"
	"fmt"
	"os"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
        	for _, arg := range os.Args[1:] {
                	s += sep + arg
                	sep = " "
        	}
        	fmt.Println(s)
	}
}
