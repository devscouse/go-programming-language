/*
Exercise 3.5
Implement a full-color Mandelbrot set using the function image.NewRGBA and the
type color.RGBA or color.YCbCr
*/

package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const xmin, ymin, xmax, ymax = -2, -2, 2, 2
	const width, height = 3840, 2160
	const upscale = 10

	const rwidth = width * upscale
	const rheight = height * upscale

	colors := palette.Plan9
	img := image.NewRGBA(image.Rect(0, 0, rwidth, rheight))
	for py := range rheight {
		y := float64(py)/rheight*(ymax-ymin) + ymin
		for px := range rwidth {
			x := float64(px)/rwidth*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z, colors))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128, colors []color.Color) color.Color {
	const iterations = 200

	var v complex128
	for n := range iterations {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colors[n]
		}
	}
	return color.Black
}
