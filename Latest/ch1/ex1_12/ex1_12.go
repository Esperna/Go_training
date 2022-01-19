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
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

var palette = []color.Color{color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, //Red
	color.RGBA{0x00, 0x80, 0x00, 0xff}, //Green
	color.RGBA{0x00, 0xff, 0x00, 0xff}, //Lime
	color.RGBA{0xff, 0xff, 0x00, 0xff}, //Yellow
	color.RGBA{0x00, 0x00, 0xff, 0xff}, //Blue
}

const (
	BlackIndex           = 0 // first color in palette: Background color
	RedIndex             = 1
	GreenIndex           = 2
	LimeIndex            = 3
	YellowIndex          = 4
	BlueIndex            = 5
	NumOfForegroundColor = 5
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var cycles int = 5
		for k, v := range r.Form {
			if k == "cycles" {
				cycles, _ = strconv.Atoi(strings.Join(v, ""))
			}
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	rand.Seed(time.Now().Unix())
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex uint8 = uint8((rand.Intn(NumOfForegroundColor)%NumOfForegroundColor + 1))
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-
