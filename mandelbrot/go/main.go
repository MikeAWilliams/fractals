package main

import (
	"fmt"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot"
)

func main() {
	fmt.Println("Starting Mandelbrot")
	inSet, iter := mandelbrot.ComputeMandelbrot(30, 0, 1)
	fmt.Printf("inset=%v iter=%v\n", inSet, iter)
	fmt.Println("Finished Mandelbrot")
}
