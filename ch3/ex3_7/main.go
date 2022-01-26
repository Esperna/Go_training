package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	width, height = 128, 128
)

func main() {
	var buf = complex(width, height)
	var z0, z1 complex128
	var count int
	img := image.NewRGBA(image.Rect(-width, -height, width, height))
	for {
		z0 = buf
		z1 = z0 - (cmplx.Pow(z0, 4)-1)/(4*cmplx.Pow(z0, 3))
		//fmt.Println(z1)
		buf = z1
		count++
		img.Set(int(real(z1)), int(imag(z1)), colorGradation(z1))
		//fmt.Printf("%d\t%d\n", int(real(z1)), int(imag(z1)))
		if cmplx.Abs(z1-z0) < 0.001 {
			break
		}
	}
	//fmt.Println(count)
	png.Encode(os.Stdout, img)
}

func colorGradation(z complex128) color.Color {
	//	fmt.Println(color.Gray{255 - uint8(255*(cmplx.Abs(z)/cmplx.Abs(complex(width, height))))})
	return color.Gray{255 - uint8(255*(cmplx.Abs(z)/cmplx.Abs(complex(width, height))))}
}
