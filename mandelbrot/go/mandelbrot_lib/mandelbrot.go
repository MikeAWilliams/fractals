package mandelbrot_lib

type Point struct {
	Real, Imaginary float64
}

type Pixel struct {
	X, Y int
}

type MandelbrotSetResult struct {
	IsIn          bool
	Iterations    int
	MaxIterations int
}

type CombinedPoint struct {
	CoordinateComplex Point
	CoordinateImage   Pixel
}

type MandelbrotPointData struct {
	Coordinates CombinedPoint
	Result      MandelbrotSetResult
}

func ComputeMandelbrot(maxIterations int, point CombinedPoint) MandelbrotPointData {
	var x2 float64
	var y2 float64
	var iteration int
	var x float64
	var y float64

	for ; iteration < maxIterations && x2+y2 < 4; iteration++ {
		y = 2*x*y + point.CoordinateComplex.Imaginary
		x = x2 - y2 + point.CoordinateComplex.Real
		x2 = x * x
		y2 = y * y
	}
	setResult := MandelbrotSetResult{maxIterations == iteration, iteration, maxIterations}
	result := MandelbrotPointData{point, setResult}
	return result
}
