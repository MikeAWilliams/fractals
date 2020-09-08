package mandelbrot_lib

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

func CreateColorMandelbrotNoPipeSingle(params Parameters, darkColor Color, lightColor Color, fileName string) {
	result := image.NewNRGBA(image.Rect(0, 0, params.MaxPixel.X, params.MaxPixel.Y))
	interpolator := GetColorInterpolator(darkColor, lightColor)

	startTime := time.Now()
	realPixelSize := (params.MaxPoint.Real - params.MinPoint.Real) / float64(params.MaxPixel.X-1)
	imaginaryPixelSize := (params.MaxPoint.Imaginary - params.MinPoint.Imaginary) / float64(params.MaxPixel.Y-1)

	for yPixel := 0; yPixel < params.MaxPixel.Y; yPixel++ {
		imaginary := params.MinPoint.Imaginary + float64(yPixel)*imaginaryPixelSize
		for xPixel := 0; xPixel < params.MaxPixel.X; xPixel++ {
			real := params.MinPoint.Real + float64(xPixel)*realPixelSize
			point := CombinedPoint{Point{real, imaginary}, Pixel{xPixel, yPixel}}
			input := MandelbrotInput{point, params.MaxIterations}

			mandPoint := ComputeMandelbrot(input)
			outPixel := ColorPixel{mandPoint.Input.Coordinates.CoordinateImage, darkColor}
			if !mandPoint.Result.IsIn {
				outPixel.Value = interpolator.Interpolate(float64(mandPoint.Result.Iterations) / float64(mandPoint.Input.MaxIterations))
			}
			result.Set(outPixel.Coordinate.X, outPixel.Coordinate.Y, color.NRGBA{outPixel.Value.R, outPixel.Value.G, outPixel.Value.B, 255})
		}
	}

	endTime := time.Now()
	fmt.Printf("The fan pipe Mandelbrot took %v", endTime.Sub(startTime))

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, result)
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
