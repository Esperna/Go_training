#! /bin/zsh

go build
./mandelbrot | ./ex01 >mandelbrot.gif
./mandelbrot | ./ex01 >mandelbrot.png
./mandelbrot | ./ex01 >mandelbrot.jpeg
open mandelbrot.gif
open mandelbrot.png
open mandelbrot.jpeg
