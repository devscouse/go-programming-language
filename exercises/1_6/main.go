// Exercise 1.6: Modify the Lissajous program to produce images in multiple
// colors by adding more values to palette and then displaying them by
// changing the third argument of SetColorIndex in some interesting ways.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xa0, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0xa0, 0xff},
}

const (
	backgroundIndex      = 0
	foregroundOneIndex   = 1
	foregroundTwoIndex   = 2
	foregroundThreeIndex = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 100
		delay   = 8
	)

	freq := rand.Float64() * 3.0

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		maxt := cycles * 2 * math.Pi
		for t := 0.0; t < maxt; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			if t < maxt/3 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), foregroundOneIndex)
			} else if t < maxt*2/3 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), foregroundTwoIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), foregroundThreeIndex)
			}

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
