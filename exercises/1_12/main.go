// Exercise 1.12: Modify the Lissajous server to read parameter values from the
// URL. For example, you might arrange it so that a URL like
// http://localhost:8000/?cycles=20 sets the number of cycles to 20 instead of
// the default 5. Use the strconv.Atoi function to convert the string param
// into an integer.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	backgroundIndex = 0
	foregroundIndex = 1
)

func lissajous(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	res := 0.001
	size := 100
	nframes := 64
	delay := 8

	for k, v := range r.URL.Query() {
		if len(v) == 0 || len(v) > 1 {
			log.Printf("Unexpected query param value for %s: %v", k, v)
			continue
		}

		value := v[0]
		if k == "cycles" {
			parsedValue, err := strconv.Atoi(value)
			if err == nil {
				cycles = parsedValue
			}
		} else if k == "res" {
			parsedValue, err := strconv.ParseFloat(value, 64)
			if err == nil {
				res = parsedValue
			}
		} else if k == "size" {
			parsedValue, err := strconv.Atoi(value)
			if err == nil {
				size = parsedValue
			}
		} else if k == "nframes" {
			parsedValue, err := strconv.Atoi(value)
			if err == nil {
				nframes = parsedValue
			}
		} else if k == "delay" {
			parsedValue, err := strconv.Atoi(value)
			if err == nil {
				delay = parsedValue
			}
		} else {
			log.Printf("Unexpected query param %s", k)
			continue
		}

	}

	freq := rand.Float64() * 3.0

	// anim is a struct of type GIF. The literal creates a struct value whose
	// LoopCount field is set to nframes; all other fields have the zero value
	// for their type.
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*float64(2)*float64(math.Pi); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				foregroundIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}
