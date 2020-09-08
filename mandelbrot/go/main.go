package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"github.com/MikeAWilliams/fractals/mandelbrot/go/mandelbrot/mandelbrot_lib"
)

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile to `file`")

func main() {
	params := mandelbrot_lib.Parameters{30, mandelbrot_lib.Point{-2.0, -1.0}, mandelbrot_lib.Point{1.0, 1.0}, mandelbrot_lib.Pixel{1920, 1080}}
	inSetColor := mandelbrot_lib.Color{0, 0, 0}
	outColor := mandelbrot_lib.Color{255, 255, 255}
	fileName := "out.png"

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	mandelbrot_lib.CreateColorMandelbrotNoPipeGoroutines(params, inSetColor, outColor, fileName)
}
