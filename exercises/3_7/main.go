/*
Exercise 3.7
Another simple fractal uses Newton's method to find complex solutions to a
function such as z^4 - 1 = 0. Shade each starting point by the number of
iterations required to get close to one of the four roots. Color each point by
the root it approaches.
*/

package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"math/rand"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	minDistToRoot          = 1e-4
	iterations             = 200
)

func pixelPosToGlobalPos(px float64, py float64, width float64, height float64) (float64, float64) {
	x := px/width*(xmax-xmin) + xmin
	y := py/height*(ymax-ymin) + ymin
	return x, y
}

func randomColor() color.RGBA {
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))
	return color.RGBA{r, g, b, 255}
}

func newtonsMethod(z complex128, rootPoints [4]complex128, rootColors [4]color.RGBA) color.Color {
	for n := range iterations {
		for i, rootPoint := range rootPoints {
			if cmplx.Abs(z-rootPoint) < minDistToRoot {
				color := rootColors[i]
				color.A = 255 - uint8(255*n/iterations)
				return color
			}
		}
		if cmplx.Abs(z) < 1e-10 {
			z += complex(1e-10, 1e-10)
		}
		z = (3*cmplx.Pow(z, 4) + 1) / (4 * cmplx.Pow(z, 3))
	}
	return color.White
}

func main() {
	const width, height = 1000, 1000
	const upscale = 10

	const rwidth = width * upscale
	const rheight = height * upscale

	rootPoints := [4]complex128{1 + 0i, -1 + 0i, 0 + 1i, 0 - 1i}
	rootColors := [4]color.RGBA{randomColor(), randomColor(), randomColor(), randomColor()}

	img := image.NewRGBA(image.Rect(0, 0, rwidth, rheight))
	for py := range rheight {
		for px := range rwidth {
			x, y := pixelPosToGlobalPos(float64(px), float64(py), rwidth, rheight)
			z := complex(x, y)
			color := newtonsMethod(z, rootPoints, rootColors)
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}
