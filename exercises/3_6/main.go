/*
Exercise 3.6
Supersampling is a technique used to reduce the effects of pixellation by computing
the color value at several points within each pixel and taking an average. The simplest
method is to divide each pixel into 4 sub-pixels.
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

const xmin, ymin, xmax, ymax = -2, -2, 2, 2

func pixelPosToGlobalPos(px float64, py float64, width float64, height float64) (float64, float64) {
	x := px/width*(xmax-xmin) + xmin
	y := py/height*(ymax-ymin) + ymin
	return x, y
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

func mandelbrotFromPixelPos(px float64, py float64, width float64, height float64, colors []color.Color) color.Color {
	x, y := pixelPosToGlobalPos(px, py, width, height)
	z := complex(x, y)
	return mandelbrot(z, colors)
}

func averageColor(colors []color.Color) color.Color {
	var r, g, b, a uint32
	count := uint32(len(colors))

	for _, c := range colors {
		cr, cg, cb, ca := c.RGBA()
		r += cr
		g += cg
		b += cb
		a += ca
	}

	return color.RGBA{
		uint8(r / count),
		uint8(g / count),
		uint8(b / count),
		uint8(a / count),
	}
}

func main() {
	const width, height = 3840, 2160
	const upscale = 1

	const rwidth = width * upscale
	const rheight = height * upscale
	const spGridSize = 2
	const spCount = spGridSize * spGridSize
	const spOffset = 0.5 / spCount
	const spSpacing = 1 / spCount

	spColors := make([]color.Color, spCount)
	colors := palette.Plan9

	img := image.NewRGBA(image.Rect(0, 0, rwidth, rheight))
	for py := range rheight {
		for px := range rwidth {
			i := 0

			for spRow := range spGridSize {
				spy := float64(py) + spOffset + float64(spRow)*spSpacing
				for spCol := range spGridSize {
					spx := float64(px) + spOffset + float64(spCol)*spSpacing
					spColors[i] = mandelbrotFromPixelPos(spx, spy, rwidth, rheight, colors)
					i++
				}
			}
			img.Set(px, py, averageColor(spColors))
		}
	}

	png.Encode(os.Stdout, img)
}
