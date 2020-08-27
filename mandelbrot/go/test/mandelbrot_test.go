package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
	"github.com/stretchr/testify/require"
)

func TestComputeMandelbrot(t *testing.T) {

	inSet, _ := mandelbrot_lib.ComputeMandelbrot(30, 0, 1)
	require.True(t, inSet)

	insetNeg10, _ := mandelbrot_lib.ComputeMandelbrot(30, -1, 0)
	require.True(t, insetNeg10)

	inset05, _ := mandelbrot_lib.ComputeMandelbrot(30, 0, 0.5)
	require.True(t, inset05)

	inset10, iter10 := mandelbrot_lib.ComputeMandelbrot(30, 1, 0)
	require.False(t, inset10)
	require.Equal(t, 2, iter10)

	inset02, iter02 := mandelbrot_lib.ComputeMandelbrot(30, 0, 2)
	require.False(t, inset02)
	require.Equal(t, 1, iter02)

	inset55, iter55 := mandelbrot_lib.ComputeMandelbrot(30, 0.5, 0.5)
	require.False(t, inset55)
	require.Equal(t, 5, iter55)
}
