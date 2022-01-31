// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	var colorValueR uint8 = 0xff
	var colorValueG uint8 = 0xff
	var colorValueB uint8 = 0xff
	prevLocalMaxI := 0
	prevLocalMaxJ := 0
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			var color string
			if i == 0 && j == 0 {
				color = "#ff0000"
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
				prevLocalMaxI = i
				prevLocalMaxJ = j
			} else if isValidCorner(i+1, j) && isValidCorner(i, j) && isValidCorner(i, j+1) && isValidCorner(i+1, j+1) {
				if isLocalMax(i, j) {
					color = "#ff0000"
					prevLocalMaxI = i
					prevLocalMaxJ = j
				} else if isLocalMin(i, j) {
					color = "#0000ff"
				} else {
					x := xyrange * (float64(i)/cells - 0.5)
					y := xyrange * (float64(j)/cells - 0.5)
					prevLocalMaxX := xyrange * (float64(prevLocalMaxI)/cells - 0.5)
					prevLocalMaxY := xyrange * (float64(prevLocalMaxJ)/cells - 0.5)
					colorValueR = uint8(f(x, y) / f(prevLocalMaxX, prevLocalMaxY) * 0xFF)
					colorValueB = uint8(0xFF * (1 - f(x, y)/f(prevLocalMaxX, prevLocalMaxY)))
					hexStrColorValR := fmt.Sprintf("%x", colorValueR)
					hexStrColorValG := fmt.Sprintf("%x", colorValueG)
					hexStrColorValB := fmt.Sprintf("%x", colorValueB)
					fmt.Println(hexStrColorValR, hexStrColorValG, hexStrColorValB)
					color = "#" + hexStrColorValR + hexStrColorValG + hexStrColorValB
					color = "#00ff00"
				}
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func isValidCorner(i, j int) bool {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	return !math.IsNaN(f(x, y))
}

func isHeightEnough(i, j int) bool {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	if f(x, y) > .003 {
		return true
	}
	return false
}

//f(x,y) - f(x-cos,y-sin)
func isLocalMax(i, j int) bool {
	x1 := xyrange * (float64(i)/cells - 0.5)
	y1 := xyrange * (float64(j)/cells - 0.5)
	cos := x1 / math.Sqrt(x1*x1+y1*y1)
	sin := y1 / math.Sqrt(x1*x1+y1*y1)
	x0 := xyrange * ((float64(i)-cos)/cells - 0.5)
	y0 := xyrange * ((float64(j)-sin)/cells - 0.5)
	x2 := xyrange * ((float64(i)+cos)/cells - 0.5)
	y2 := xyrange * ((float64(j)+sin)/cells - 0.5)

	a := f(x1, y1) - f(x0, y0)
	b := f(x2, y2) - f(x1, y1)

	if (a*b <= 0) && (a >= 0) {
		return true
	}
	return false
}

func isLocalMin(i, j int) bool {
	x1 := xyrange * (float64(i)/cells - 0.5)
	y1 := xyrange * (float64(j)/cells - 0.5)
	cos := x1 / math.Sqrt(x1*x1+y1*y1)
	sin := y1 / math.Sqrt(x1*x1+y1*y1)
	x0 := xyrange * ((float64(i)-cos)/cells - 0.5)
	y0 := xyrange * ((float64(j)-sin)/cells - 0.5)
	x2 := xyrange * ((float64(i)+cos)/cells - 0.5)
	y2 := xyrange * ((float64(j)+sin)/cells - 0.5)

	a := f(x1, y1) - f(x0, y0)
	b := f(x2, y2) - f(x1, y1)
	if (a*b <= 0) && (b > 0) {
		return true
	}
	return false
}

//!-
