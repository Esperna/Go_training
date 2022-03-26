// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}(px, py)
		}
	}

	wg.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	secs := time.Since(start).Seconds()
	fmt.Printf("\n%.2fs\n", secs)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func drawMandelbrotWithoutGoRoutine() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(ioutil.Discard, img) // NOTE: ignoring errors
}

func drawMandelbrotWithGoRoutine() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(ioutil.Discard, img) // NOTE: ignoring errors
}

func drawMandelbrotWithGoRoutineWithBuffer(numOfBuf int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ch := make(chan struct{}, numOfBuf)

	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot2(ch, z))
			}(px, py)
		}
	}
	wg.Wait()
	png.Encode(ioutil.Discard, img) // NOTE: ignoring errors
}

func mandelbrot2(ch chan struct{}, z complex128) color.Color {
	ch <- struct{}{}
	defer func() { <-ch }()
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
