package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
	"github.com/stretchr/testify/require"
)

func TestComputeMandelbrot(t *testing.T) {

	inPoint := mandelbrot_lib.CombinedPoint{mandelbrot_lib.Point{0.0, 1.0}, mandelbrot_lib.Pixel{2, 3}}
	result01 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.True(t, result01.Result.IsIn)
	require.Equal(t, 30, result01.Result.MaxIterations)
	require.Equal(t, 0.0, result01.Coordinates.CoordinateComplex.Real)
	require.Equal(t, 1.0, result01.Coordinates.CoordinateComplex.Imaginary)
	require.Equal(t, 2, result01.Coordinates.CoordinateImage.X)
	require.Equal(t, 3, result01.Coordinates.CoordinateImage.Y)

	inPoint.CoordinateComplex.Real = -1
	inPoint.CoordinateComplex.Imaginary = 0
	resultNeg10 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.True(t, resultNeg10.Result.IsIn)

	inPoint.CoordinateComplex.Real = 0.0
	inPoint.CoordinateComplex.Imaginary = 0.5
	result05 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.True(t, result05.Result.IsIn)

	inPoint.CoordinateComplex.Real = 1.0
	inPoint.CoordinateComplex.Imaginary = 0.0
	result10 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.False(t, result10.Result.IsIn)
	require.Equal(t, 2, result10.Result.Iterations)

	inPoint.CoordinateComplex.Real = 0.0
	inPoint.CoordinateComplex.Imaginary = 2.0
	result02 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.False(t, result02.Result.IsIn)
	require.Equal(t, 1, result02.Result.Iterations)

	inPoint.CoordinateComplex.Real = 0.5
	inPoint.CoordinateComplex.Imaginary = 0.5
	result55 := mandelbrot_lib.ComputeMandelbrot(30, inPoint)
	require.False(t, result55.Result.IsIn)
	require.Equal(t, 5, result55.Result.Iterations)
}

func TestPointGeneratorNoDone(t *testing.T) {
	params := mandelbrot_lib.Parameters{30, mandelbrot_lib.Point{-2.0, -1.0}, mandelbrot_lib.Point{1.0, 1.0}, mandelbrot_lib.Pixel{300, 300}}

	done := make(chan struct{})
	defer close(done)

	pointStream := mandelbrot_lib.PointGenerator(done, params)

	count := 0
	for point := range pointStream {
		count++
		if point.CoordinateComplex.Real > 500 {
			panic("oh man")
		}
	}

	require.Equal(t, 90000, count)

}
