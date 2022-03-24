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

func BenchmarkDrawMandelbrotWithGoRoutineWith1Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer(1)
	}
}

func BenchmarkDrawMandelbrotWithGoRoutineWith2Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer(2)
	}
}

func BenchmarkDrawMandelbrotWithGoRoutineWith4Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer(4)
	}
}

func BenchmarkDrawMandelbrotWithGoRoutineWith8Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer(8)
	}
}

func BenchmarkDrawMandelbrotWithGoRoutineWith16Buffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		drawMandelbrotWithGoRoutineWithBuffer(16)
	}
}
