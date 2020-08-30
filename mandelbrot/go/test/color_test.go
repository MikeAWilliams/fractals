package testMandelbrot

import (
	"testing"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
	"github.com/stretchr/testify/require"
)

func TestInterpolate(t *testing.T) {
	darkerColor := mandelbrot_lib.Color{1, 51, 75}
	lighterColor := mandelbrot_lib.Color{255, 255, 255}

	testObject := mandelbrot_lib.GetColorInterpolator(darkerColor, lighterColor)

	middleColor := testObject.Interpolate(0.5)

	require.Equal(t, byte(128), middleColor.R)
	require.Equal(t, byte(153), middleColor.G)
	require.Equal(t, byte(165), middleColor.B)
}
