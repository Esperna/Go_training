// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(&buf, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	io.WriteString(w, buf.String())
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if isValidCorner(i+1, j) && isValidCorner(i, j) && isValidCorner(i, j+1) && isValidCorner(i+1, j+1) {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				var color string
				if isHeightEnough(i, j) {
					color = "#ff0000"
				} else {
					color = "#0000ff"
				}
				var buf2 bytes.Buffer
				fmt.Fprintf(&buf2, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
				io.WriteString(w, buf2.String())
			}

		}
	}
	io.WriteString(w, "</svg>\n")
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

//!-handler