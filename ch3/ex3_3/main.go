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
	maxZ, minZ := findMinMax()
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
			if isValidCorner(i+1, j) && isValidCorner(i, j) && isValidCorner(i, j+1) && isValidCorner(i+1, j+1) {
				x := xyrange * (float64(i)/cells - 0.5)
				y := xyrange * (float64(j)/cells - 0.5)
				r := float64((f(x, y) - minZ) / (maxZ - minZ))
				var colorValueR, colorValueG, colorValueB uint8
				colorValueR = uint8(255.0 * r)
				colorValueB = uint8(255.0 * (1.0 - r))
				var hexStrColorValR, hexStrColorValG, hexStrColorValB string
				if colorValueR < 0x10 {
					hexStrColorValR = fmt.Sprintf("0%x", colorValueR)
				} else {
					hexStrColorValR = fmt.Sprintf("%x", colorValueR)
				}
				if colorValueG < 0x10 {
					hexStrColorValG = fmt.Sprintf("0%x", colorValueG)
				} else {
					hexStrColorValG = fmt.Sprintf("%x", colorValueG)
				}
				if colorValueB < 0x10 {
					hexStrColorValB = fmt.Sprintf("0%x", colorValueB)
				} else {
					hexStrColorValB = fmt.Sprintf("%x", colorValueB)
				}
				//					fmt.Println(hexStrColorValR, hexStrColorValG, hexStrColorValB)
				color = "#" + hexStrColorValR + hexStrColorValG + hexStrColorValB
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}

	}
	fmt.Println("</svg>")

}

type Point struct {
	x          float64
	y          float64
	z          float64
	isLocalMax bool
}

func findMinMax() (float64, float64) {
	max := -math.MaxFloat64
	min := math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			if isValidCorner(i+1, j) && isValidCorner(i, j) && isValidCorner(i, j+1) && isValidCorner(i+1, j+1) {
				z := f(x, y)
				if z > max {
					max = z
				}
				if z < min {
					min = z
				}
			}
		}
	}
	//fmt.Printf("Max\t%g\tMin\t%g", max, min)
	return max, min
}

func findLocalMinMax() {
	localQ := make([]Point, 0)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			if isValidCorner(i+1, j) && isValidCorner(i, j) && isValidCorner(i, j+1) && isValidCorner(i+1, j+1) {
				if isLocalMax(i, j) {
					localQ = append(localQ, Point{x, y, f(x, y), true})
				} else if isLocalMin(i, j) {
					localQ = append(localQ, Point{x, y, f(x, y), false})
				} else {

				}
			}
		}
	}
	fmt.Printf("Number Of local:\t%d\n", len(localQ))
	for _, v := range localQ {
		fmt.Printf("local \t%v\n", v)
	}

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
	//a := f(x1, y1) - f(x1, y0)
	//b := f(x2, y2) - f(x2, y1)

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
	//a := f(x1, y1) - f(x1, y0)
	//b := f(x2, y2) - f(x2, y1)

	if (a*b <= 0) && (b > 0) {
		return true
	}
	return false
}

//!-
