package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
	"github.com/stretchr/testify/require"
)

func TestComputeMandelbrot(t *testing.T) {

	inSet, _ := mandelbrot_lib.ComputeMandelbrot(30, 0, 1)
	require.True(t, inSet)
}
