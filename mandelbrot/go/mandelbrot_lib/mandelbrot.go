package mandelbrot_lib

type Point struct {
	Real, Imaginary float64
}

type Pixel struct {
	X, Y int
}

type MandelbrotSetResult struct {
	IsIn       bool
	Iterations int
}

type CombinedPoint struct {
	CoordinateComplex Point
	CoordinateImage   Pixel
}

type MandelbrotInput struct {
	Coordinates   CombinedPoint
	MaxIterations int
}

type MandelbrotPointData struct {
	Input  MandelbrotInput
	Result MandelbrotSetResult
}

type Parameters struct {
	MaxIterations int
	MinPoint      Point
	MaxPoint      Point
	MaxPixel      Pixel
}

func ComputeMandelbrot(input MandelbrotInput) MandelbrotPointData {
	var x2 float64
	var y2 float64
	var iteration int
	var x float64
	var y float64

	for ; iteration < input.MaxIterations && x2+y2 < 4; iteration++ {
		y = 2*x*y + input.Coordinates.CoordinateComplex.Imaginary
		x = x2 - y2 + input.Coordinates.CoordinateComplex.Real
		x2 = x * x
		y2 = y * y
	}
	setResult := MandelbrotSetResult{input.MaxIterations == iteration, iteration}
	result := MandelbrotPointData{input, setResult}
	return result
}
