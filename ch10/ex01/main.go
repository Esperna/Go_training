// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif" // register PNG decoder
	"image/jpeg"
	"image/png"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

var o = flag.String("o", "png", "image file format")

func main() {
	flag.Parse()
	if *o == "gif" {
		if err := toGIF(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", *o, err)
			os.Exit(1)
		}
	} else if *o == "png" {
		if err := toPNG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", *o, err)
			os.Exit(1)
		}
	} else if *o == "jpeg" {
		if err := toJPEG(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", *o, err)
			os.Exit(1)
		}

	}
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{NumColors: 256})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
