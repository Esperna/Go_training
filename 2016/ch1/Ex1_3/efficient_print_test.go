package main

import(
	"testing"
)

func BenchmarkEfficientPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientPrintln()
	}
}
