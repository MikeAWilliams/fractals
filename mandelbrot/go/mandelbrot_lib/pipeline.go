package mandelbrot_lib

import "fmt"

func PointGenerator(done <-chan struct{}, params Parameters) <-chan MandelbrotInput {
	outStream := make(chan MandelbrotInput)
	go func() {
		defer close(outStream)

		realPixelSize := (params.MaxPoint.Real - params.MinPoint.Real) / float64(params.MaxPixel.X-1)
		imaginaryPixelSize := (params.MaxPoint.Imaginary - params.MinPoint.Imaginary) / float64(params.MaxPixel.Y-1)

		for yPixel := 0; yPixel < params.MaxPixel.Y; yPixel++ {
			imaginary := params.MinPoint.Imaginary + float64(yPixel)*imaginaryPixelSize
			for xPixel := 0; xPixel < params.MaxPixel.X; xPixel++ {
				real := params.MinPoint.Real + float64(xPixel)*realPixelSize
				point := CombinedPoint{Point{real, imaginary}, Pixel{xPixel, yPixel}}
				input := MandelbrotInput{point, params.MaxIterations}
				select {
				case <-done:
					return
				case outStream <- input:
				}
			}
		}
	}()
	return outStream
}

func MandelbrotPointDataGenerator(done <-chan struct{}, points <-chan MandelbrotInput) <-chan MandelbrotPointData {
	outStream := make(chan MandelbrotPointData)
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case point, ok := <-points:
				if !ok {
					return
				}
				outStream <- ComputeMandelbrot(point)
			}
		}
	}()
	return outStream
}

type ASCIIPixel struct {
	Coordinate Pixel
	Value      byte
}

func ASCIIPointGenerator(done <-chan struct{}, points <-chan MandelbrotPointData) <-chan ASCIIPixel {
	outStream := make(chan ASCIIPixel)
	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case point, ok := <-points:
				if !ok {
					return
				}
				outPixel := ASCIIPixel{point.Input.Coordinates.CoordinateImage, 'x'}
				if point.Result.IsIn {
					outPixel.Value = ' '
				}
				outStream <- outPixel
			}
		}
	}()
	return outStream
}

func PrintASCIIMandelbrot(params Parameters) {
	result := [][]byte{}
	for y := 0; y < params.MaxPixel.Y; y++ {
		result = append(result, make([]byte, params.MaxPixel.X))
	}

	done := make(chan struct{})
	defer close(done)
	asciiStream := ASCIIPointGenerator(done, MandelbrotPointDataGenerator(done, PointGenerator(done, params)))

	for asciiPoint := range asciiStream {
		result[asciiPoint.Coordinate.Y][asciiPoint.Coordinate.X] = asciiPoint.Value
	}

	for _, row := range result {
		fmt.Println(string(row))
	}
}
