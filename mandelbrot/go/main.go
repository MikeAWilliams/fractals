package main

import (
	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
)

func main() {
	params := mandelbrot_lib.Parameters{30, mandelbrot_lib.Point{-2.0, -1.0}, mandelbrot_lib.Point{1.0, 1.0}, mandelbrot_lib.Pixel{70, 30}}
	mandelbrot_lib.PrintASCIIMandelbrot(params)
}
