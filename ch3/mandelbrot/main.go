// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2
	const width, height = 2160, 2160

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := range height {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := range width {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128
	for n := range iterations {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*uint8(n)}
		}
	}
	return color.Black
}
