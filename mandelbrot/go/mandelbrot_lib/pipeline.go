package mandelbrot_lib

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
	"time"
)

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

func MandelbrotPointDataCalculatorSingle(done <-chan struct{}, points <-chan MandelbrotInput) <-chan MandelbrotPointData {
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

func FanIn(done <-chan struct{}, channels ...<-chan MandelbrotPointData) <-chan MandelbrotPointData {
	var waitGroup sync.WaitGroup
	multiplexedStream := make(chan MandelbrotPointData)

	multiplexFunction := func(c <-chan MandelbrotPointData) {
		defer waitGroup.Done()
		for data := range c {
			select {
			case <-done:
				return
			case multiplexedStream <- data:
			}
		}
	}

	waitGroup.Add(len(channels))
	for _, c := range channels {
		go multiplexFunction(c)
	}

	go func() {
		waitGroup.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func MandelbrotPointDataCalculatorFan(done <-chan struct{}, points <-chan MandelbrotInput, numchanels int) <-chan MandelbrotPointData {
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

func ASCIIPointCalculator(done <-chan struct{}, points <-chan MandelbrotPointData) <-chan ASCIIPixel {
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
	asciiStream := ASCIIPointCalculator(done, MandelbrotPointDataCalculatorSingle(done, PointGenerator(done, params)))

	for asciiPoint := range asciiStream {
		result[asciiPoint.Coordinate.Y][asciiPoint.Coordinate.X] = asciiPoint.Value
	}

	for _, row := range result {
		fmt.Println(string(row))
	}
}

type ColorPixel struct {
	Coordinate Pixel
	Value      Color
}

func ColorPointCalculator(done <-chan struct{}, points <-chan MandelbrotPointData, inSetColor Color, interpolator ColorInterpolator) <-chan ColorPixel {
	outStream := make(chan ColorPixel)
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
				outPixel := ColorPixel{point.Input.Coordinates.CoordinateImage, inSetColor}
				if !point.Result.IsIn {
					outPixel.Value = interpolator.Interpolate(float64(point.Result.Iterations) / float64(point.Input.MaxIterations))
				}
				outStream <- outPixel
			}
		}
	}()
	return outStream
}

func CreateColorMandelbrotSingle(params Parameters, darkColor Color, lightColor Color, fileName string) {
	result := image.NewNRGBA(image.Rect(0, 0, params.MaxPixel.X, params.MaxPixel.Y))

	interpolator := GetColorInterpolator(darkColor, lightColor)

	done := make(chan struct{})
	defer close(done)
	colorStream := ColorPointCalculator(done, MandelbrotPointDataCalculatorSingle(done, PointGenerator(done, params)), darkColor, interpolator)

	startTime := time.Now()
	for point := range colorStream {
		result.Set(point.Coordinate.X, point.Coordinate.Y, color.NRGBA{point.Value.R, point.Value.G, point.Value.B, 255})
	}
	endTime := time.Now()
	// around 2.8 seconds on laptop
	fmt.Printf("The single pipe Mandelbrot took %v", endTime.Sub(startTime))

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

func CreateColorMandelbrotFan(params Parameters, darkColor Color, lightColor Color, fileName string) {
	result := image.NewNRGBA(image.Rect(0, 0, params.MaxPixel.X, params.MaxPixel.Y))

	interpolator := GetColorInterpolator(darkColor, lightColor)

	done := make(chan struct{})
	defer close(done)
	colorStream := ColorPointCalculator(done, MandelbrotPointDataCalculatorFan(done, PointGenerator(done, params), runtime.NumCPU()), darkColor, interpolator)

	startTime := time.Now()
	for point := range colorStream {
		result.Set(point.Coordinate.X, point.Coordinate.Y, color.NRGBA{point.Value.R, point.Value.G, point.Value.B, 255})
	}
	endTime := time.Now()
	// single is around 2.8 seconds on laptop
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
