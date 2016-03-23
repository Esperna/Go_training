// Copyright ﾂｩ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var width, height = 800, 600                // canvas size in pixels
var xyrange = 30                            // axis ranges (-xyrange..+xyrange)
var xyscale = width / 2 / xyrange           // pixels per x or y unit
var zscale = float64(height) * float64(0.4) // pixels per z unit
var color = "grey"

const (
	cells = 100         // number of grid cells
	angle = math.Pi / 6 // angle of x, y axes (=30ﾂｰ)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30ﾂｰ), cos(30ﾂｰ)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			for k, v := range r.Form {
				if k == "height" {
					height, _ = strconv.Atoi(strings.Join(v, ""))
				} else if k == "width" {
					width, _ = strconv.Atoi(strings.Join(v, ""))
				} else if k == "color" {
					color = v[0]
				}
			}
			w.Header().Set("Content-Type", "image/svg+xml")
			drawSVG(w, r)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	http.HandleFunc("/draw", drawSVG)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func drawSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"viewBox='0 0 %d %d' "+
		"x='720' y='375' width='%d' height='%d'>", color, width, height, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := float64(xyrange) * (float64(i)/cells - 0.5)
	y := float64(xyrange) * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + float64((x-y))*cos30*float64(xyscale)
	sy := float64(height/2) + float64((x+y))*sin30*float64(xyscale) - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
