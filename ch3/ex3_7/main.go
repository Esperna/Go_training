package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	width, height = 128, 128
)

func main() {
	img := image.NewRGBA(image.Rect(-width, -height, width, height))
	drawSolution(img, complex(width, height))
	png.Encode(os.Stdout, img)
}

func colorGradation(z complex128) color.Color {
	gradPercentage := 1.0 - math.Log(cmplx.Abs(z))/math.Log(cmplx.Abs(complex(width, height)))
	//	fmt.Println(gradPercentage)
	return color.RGBA{50, 100, uint8(255 * gradPercentage), 255}
}

func drawSolution(img *image.RGBA, startPoint complex128) {
	var count int
	var buf = startPoint
	var z0, z1 complex128
	for {
		z0 = buf
		z1 = z0 - (cmplx.Pow(z0, 4)-1)/(4*cmplx.Pow(z0, 3))
		buf = z1
		count++
		img.Set(int(real(z1)), int(imag(z1)), colorGradation(z1))
		if cmplx.Abs(z1-z0) < 0.001 {
			break
		}
	}
}
