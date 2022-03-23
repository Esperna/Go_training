package main

import "testing"

func BenchmarkDrawMandelbrotWithoutGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithoutGoRoutine()
	}
}

func BenchmarkDrawMandelbrotWithGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutine()
	}
}

func BenchmarkDrawMandelbrotWithGoRoutineWithBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer()
	}
}
