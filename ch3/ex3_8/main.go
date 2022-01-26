// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//drawFractalComplex64(img)
	//drawFractalComplex128(img)
	//drawFractalBigFloat(img)
	drawFractal(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func drawFractal(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			//img.Set(px, py, sqrt(z))
			img.Set(px, py, acos(z))
		}
	}
}

func drawFractalComplex64(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
}

func drawFractalComplex128(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
}

func drawFractalBigFloat(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotBigFloatComplex128(x, y))
		}
	}
}

func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			//			return color.RGBA{255 - contrast*n, 230 - contrast*n, 200 - contrast*n, 255}
			return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
		}
	}
	return color.YCbCr{255, 255, 255}
}

func mandelbrotComplex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//			return color.RGBA{255 - contrast*n, 230 - contrast*n, 200 - contrast*n, 255}
			return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
		}
	}
	return color.YCbCr{255, 255, 255}
}

func mandelbrotBigFloatComplex128(x, y float64) color.Color {
	// pending because cannot implement
	/*	const iterations = 200
		const contrast = 15

		var u, v float64
		for n := uint8(0); n < iterations; n++ {
			realV := u*u - v*v + x
			imagV := 2*u*v + y
			normV := big.Sqrt(big.NewFloat(realV*realV + imagV*imagV))
			if normV > 2.0 {
				//			return color.RGBA{255 - contrast*n, 230 - contrast*n, 200 - contrast*n, 255}
				return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
			}
		}
	*/
	return color.YCbCr{255, 255, 255}
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
