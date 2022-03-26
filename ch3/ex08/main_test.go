package main

import (
	"image"
	"testing"
)

func BenchmarkDrawFractal(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		DrawFractal(img)
	}
}

func BenchmarkDrawFractalComplex64(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		DrawFractalComplex64(img)
	}
}

func BenchmarkDrawFractalComplex128(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		DrawFractalComplex128(img)
	}
}

func BenchmarkDrawFractalBigFloat(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		DrawFractalBigFloat(img)
	}
}
