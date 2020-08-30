package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
)

func BenchmarkASCIIMandelbrotPipeline(b *testing.B) {
	params := mandelbrot_lib.Parameters{30, mandelbrot_lib.Point{-2.0, -1.0}, mandelbrot_lib.Point{1.0, 1.0}, mandelbrot_lib.Pixel{70, 30}}
	result := [][]byte{}
	for y := 0; y < params.MaxPixel.Y; y++ {
		result = append(result, make([]byte, params.MaxPixel.X))
	}

	done := make(chan struct{})
	defer close(done)
	asciiStream := mandelbrot_lib.ASCIIPointGenerator(done, mandelbrot_lib.MandelbrotPointDataGenerator(done, mandelbrot_lib.PointGenerator(done, params)))

	for asciiPoint := range asciiStream {
		result[asciiPoint.Coordinate.Y][asciiPoint.Coordinate.X] = asciiPoint.Value
	}
}
