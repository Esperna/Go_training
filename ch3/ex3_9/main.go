// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var x, y int
		var scale float64 = 1.0

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			if k == "x" {
				x, _ = strconv.Atoi(v[0])
			} else if k == "y" {
				y, _ = strconv.Atoi(v[0])
			} else if k == "scale" {
				scale, _ = strconv.ParseFloat(v[0], 64)
			}
			fmt.Println(k, v)
		}
		fractal(w, x, y, scale)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func fractal(out io.Writer, x_init, y_init int, scale float64) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(float64(x_init)+x/scale, float64(y_init)+y/scale)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
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

//!-
