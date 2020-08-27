package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/tree/master/mandelbrot/go/mandelbrot"
	"github.com/stretchr/testify/require"
)

func TEST_SOMETHING(t *testing.T) {

	inSet, iterations := mandelbrot.ComputeMandelbrot(30, 0, 1)
	require.True(inSet)
}
