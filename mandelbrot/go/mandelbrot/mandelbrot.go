package mandelbrot

func ComputeMandelbrot(maxIterations int, c0Real, c0Im float64) (bool, int) {
	x2 := 0.0
	y2 := 0.0
    iteration := 0
    x := 0.0
    y := 0.0

    for(; iteration < maxIterations && x2 + y2 < 4; iteration++) {
        y = 2 * x * y + c0Im;
        x = x2 - y2 + c0Real;
        x2 = x * x;
        y2 = y * y;
	}
	if maxIterations == iteration {
		return true, 0
	}
	return false, iteration
}
