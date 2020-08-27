package testMandelbrot

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TEST_SOMETHING(t *testing.T) {

	inSet, _ := mandelbrot.ComputeMandelbrot(30, 0, 1)
	require.True(t, inSet)
}
