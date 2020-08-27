package main

import (
	"fmt"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
)

func main() {
	fmt.Println("Starting Mandelbrot")
	inSet, iter := mandelbrot_lib.ComputeMandelbrot(30, 0, 1)
	fmt.Printf("inset=%v iter=%v\n", inSet, iter)
	fmt.Println("Finished Mandelbrot")
}
