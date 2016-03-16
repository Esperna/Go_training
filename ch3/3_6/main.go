// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
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
	"math/cmplx"
	"os"
)

func main() {
	inputFile, err := os.Open("test.png")

	if nil != err {
		fmt.Println(err)
		return
	}
	inputImage, _, err := image.Decode(inputFile)

	if nil != err {
		fmt.Println(err)
	}

	defer inputFile.Close()

	outputFile, err := os.Create("test_out.png")
	if nil != err {
		fmt.Println(err)
	}
	outputImage := superSampling(inputImage)
	err = png.Encode(outputFile, outputImage)

	if nil != err {
		fmt.Println(err)
	}

	defer outputFile.Close()
}

func superSampling(inputImage image.Image) image.Image {
	rect := inputImage.Bounds()
	width := rect.Size().X
	height := rect.Size().Y
	rect2 := image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X-1, rect.Max.Y-1)
	rgba := image.NewRGBA(rect2)

	for x := 0; x < width-1; x++ {
		for y := 0; y < height-1; y++ {
			var col color.RGBA
			// 座標(x,y)のR, G, B, α の値を取得
			r00, g00, b00, a00 := inputImage.At(x, y).RGBA()
			r01, g01, b01, a01 := inputImage.At(x, y+1).RGBA()
			r10, g10, b10, a10 := inputImage.At(x+1, y).RGBA()
			r11, g11, b11, a11 := inputImage.At(x+1, y+1).RGBA()
			col.R = uint8((uint(uint8(r00)) + uint(uint8(r01)) + uint(uint8(r10)) + uint(uint8(r11))) / 4)
			col.G = uint8((uint(uint8(g00)) + uint(uint8(g01)) + uint(uint8(g10)) + uint(uint8(g11))) / 4)
			col.B = uint8((uint(uint8(b00)) + uint(uint8(b01)) + uint(uint8(b10)) + uint(uint8(b11))) / 4)
			col.A = uint8((uint(uint8(a00)) + uint(uint8(a01)) + uint(uint8(a10)) + uint(uint8(a11))) / 4)
			rgba.Set(x, y, col)
		}
	}

	return rgba.SubImage(rect)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return sqrt(z)
			//return color.Gray{255 - contrast*n}
		}
	}
	return acos(z)
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
