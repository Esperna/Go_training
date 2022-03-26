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
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//DrawFractalComplex64(img)
	//DrawFractalComplex128(img)
	//DrawFractalBigFloat(img)
	DrawFractalBigRat(img)
	//DrawFractal(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func DrawFractal(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, acos(z))
		}
	}
}

func DrawFractalComplex64(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
}

func DrawFractalComplex128(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
}

func DrawFractalBigFloat(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin)
			img.Set(px, py, mandelbrotBigFloatComplex128(x, y))
		}
	}
}

func DrawFractalBigRat(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := new(big.Rat).Add(new(big.Rat).Mul(big.NewRat(int64(py), height), new(big.Rat).Sub(big.NewRat(ymax, 1), big.NewRat(ymin, 1))), big.NewRat(ymin, 1))
		for px := 0; px < width; px++ {
			x := new(big.Rat).Add(new(big.Rat).Mul(big.NewRat(int64(px), width), new(big.Rat).Sub(big.NewRat(xmax, 1), big.NewRat(xmin, 1))), big.NewRat(xmin, 1))
			img.Set(px, py, mandelbrotBigRatComplex128(x, y))
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
			return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
		}
	}
	return color.YCbCr{255, 255, 255}
}

func mandelbrotBigFloatComplex128(x, y *big.Float) color.Color {
	const iterations = 200
	const contrast = 15
	var u = big.NewFloat(0)
	var v = big.NewFloat(0)
	for n := uint8(0); n < iterations; n++ {
		realV := new(big.Float).Add(new(big.Float).Sub(new(big.Float).Mul(u, u), new(big.Float).Mul(v, v)), x)
		imagV := new(big.Float).Add(new(big.Float).Mul(big.NewFloat(2), new(big.Float).Mul(u, v)), y)
		normV := new(big.Float).Sqrt(new(big.Float).Add(new(big.Float).Mul(realV, realV), new(big.Float).Mul(imagV, imagV)))
		u = realV
		v = imagV
		if normV.Cmp(big.NewFloat(2)) > 0 {
			return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
		}
	}
	return color.YCbCr{255, 255, 255}
}

func mandelbrotBigRatComplex128(x, y *big.Rat) color.Color {
	const iterations = 200
	const contrast = 15
	var u = big.NewRat(0, 1)
	var v = big.NewRat(0, 1)
	for n := uint8(0); n < iterations; n++ {
		realV := new(big.Rat).Add(new(big.Rat).Sub(new(big.Rat).Mul(u, u), new(big.Rat).Mul(v, v)), x)
		imagV := new(big.Rat).Add(new(big.Rat).Mul(big.NewRat(2, 1), new(big.Rat).Mul(u, v)), y)
		normV := new(big.Rat).Add(new(big.Rat).Mul(realV, realV), new(big.Rat).Mul(imagV, imagV))
		u = realV
		v = imagV
		if normV.Cmp(big.NewRat(4, 1)) > 0 {
			return color.YCbCr{255 - contrast*n, 230 - contrast*n, 200 - contrast*n}
		}
	}
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

//!-
