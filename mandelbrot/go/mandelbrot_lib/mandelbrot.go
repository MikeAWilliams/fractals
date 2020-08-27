package mandelbrot_lib

func ComputeMandelbrot(maxIterations int, c0Real, c0Im float64) (bool, int) {
	var x2 float64
	var y2 float64
	var iteration int
	var x float64
	var y float64

	for ; iteration < maxIterations && x2+y2 < 4; iteration++ {
		y = 2*x*y + c0Im
		x = x2 - y2 + c0Real
		x2 = x * x
		y2 = y * y
	}
	if maxIterations == iteration {
		return true, 0
	}
	return false, iteration
}
