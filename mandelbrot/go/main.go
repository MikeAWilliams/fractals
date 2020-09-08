package main

import (
	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
)

func main() {
	params := mandelbrot_lib.Parameters{30, mandelbrot_lib.Point{-2.0, -1.0}, mandelbrot_lib.Point{1.0, 1.0}, mandelbrot_lib.Pixel{1920, 1080}}
	inSetColor := mandelbrot_lib.Color{0, 0, 0}
	outColor := mandelbrot_lib.Color{255, 255, 255}
	fileName := "out.png"

	mandelbrot_lib.CreateColorMandelbrotNoPipeGoroutines(params, inSetColor, outColor, fileName)
}
