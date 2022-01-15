package main

import(
	"testing"
)

func BenchmarkInefficientPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InefficientPrintln()
	}
}

