// Lissajous generates GIF animations of random Lissajous figures
package main

// When importing a package whose path has multiple components like
// image/color, we refer to the package with a name that comes from the last
// component.
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// This expression is a composite literal, a compact notation for instantiating
// any of Go's composite types from a sequence of element values. This one is a
// a slice and gif.GIF{LoopCount: nframes} is a struct.
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0

	// anim is a struct of type GIF. The literal creates a struct value whose
	// LoopCount field is set to nframes; all other fields have the zero value
	// for their type.
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
