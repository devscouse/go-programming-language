/*
Exercise 3.9
Write a web server that renders fractals and writes the image data to the client.
Allow the client to specify the x, y, and zoom values to the HTTP request.
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	width  = 1000
	height = 1000
)

type ViewInfo struct {
	x    float64
	y    float64
	zoom float64

	xMin float64
	xMax float64
	yMin float64
	yMax float64
}

func NewViewInfo(x float64, y float64, zoom float64) ViewInfo {
	viewSize := 4.0 / zoom
	viewOffset := viewSize / 2.0
	log.Printf("DEBUG: viewSize=%f, viewOffset=%f", viewSize, viewOffset)

	xMin := x - viewOffset
	xMax := x + viewOffset
	yMin := y - viewOffset
	yMax := y + viewOffset

	return ViewInfo{
		x:    x,
		y:    y,
		zoom: zoom,
		xMin: xMin,
		xMax: xMax,
		yMin: yMin,
		yMax: yMax,
	}
}

func main() {
	http.HandleFunc("/fractal", fractal)
	log.Fatal(http.ListenAndServe("localhost:4040", nil))
}

func parseQueryParamFloat(r *http.Request, k string, d float64) (float64, error) {
	if !r.URL.Query().Has(k) {
		return d, nil
	}
	f, err := strconv.ParseFloat(r.URL.Query().Get(k), 64)
	return f, err
}

func fractal(w http.ResponseWriter, r *http.Request) {
	fractalType := r.URL.Query().Get("type")
	if fractalType == "" {
		fractalType = "mandelbrot"
	}
	viewX, err := parseQueryParamFloat(r, "x", 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing param x: %s", err), http.StatusBadRequest)
		return
	}

	viewY, err := parseQueryParamFloat(r, "y", 0)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing param y: %s", err), http.StatusBadRequest)
		return
	}

	zoom, err := parseQueryParamFloat(r, "zoom", 1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing param zoom: %s", err), http.StatusBadRequest)
		return
	}
	if zoom <= 0 {
		http.Error(w, fmt.Sprintf("Zoom must be a positive value: got %f", zoom), http.StatusBadRequest)
		return
	}

	switch fractalType {

	case "mandelbrot":
		mandelbrotRender(NewViewInfo(viewX, viewY, zoom), w)

	default:
		http.Error(w, fmt.Sprintf("Unsupported value for type: got %s", fractalType), http.StatusBadRequest)
	}
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

func mandelbrotRender(viewInfo ViewInfo, w http.ResponseWriter) {
	log.Printf("rendering mandelbrot with view: %+v", viewInfo)
	colors := palette.Plan9
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := range height {
		y := float64(py)/height*(viewInfo.yMax-viewInfo.yMin) + viewInfo.yMin
		for px := range width {
			x := float64(px)/width*(viewInfo.xMax-viewInfo.xMin) + viewInfo.xMin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z, colors))
		}
	}
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
