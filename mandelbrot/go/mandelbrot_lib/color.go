package mandelbrot_lib

type Color struct {
	R, G, B byte
}

type ColorInterpolator struct {
	darkColor  Color
	colorDelta Color
}

func GetColorInterpolator(darkerColor, lighterColor Color) ColorInterpolator {
	return ColorInterpolator{darkerColor,
		Color{lighterColor.R - darkerColor.R,
			lighterColor.G - darkerColor.G,
			lighterColor.B - darkerColor.B}}
}

func (interp ColorInterpolator) Interpolate(fraction float64) Color {
	return Color{interp.darkColor.R + byte(fraction*float64(interp.colorDelta.R)),
		interp.darkColor.G + byte(fraction*float64(interp.colorDelta.G)),
		interp.darkColor.B + byte(fraction*float64(interp.colorDelta.B))}
}
