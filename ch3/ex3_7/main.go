package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	const (
		width, height = 1024, 1024
	)
	var buf = complex(width, height)
	var z0, z1 complex128
	var count int
	for {
		z0 = buf
		z1 = z0 - (cmplx.Pow(z0, 4)-1)/(4*cmplx.Pow(z0, 3))
		fmt.Println(z1)
		buf = z1
		count++
		//img.Set(real(z1), imag(z1), color.Black)
		if cmplx.Abs(z1-z0) < 0.001 {
			break
		}
	}
	fmt.Println(count)
}
