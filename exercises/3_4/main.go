/*
Exercise 3.4:
Following the appraoch of the Lissajous example in Section 1.7, construct a web
server that computes surfaces and writes SVG data to the client. The server
must set the Content-Type header like this:

w.Header().Set("Content-Type", "image/svg+xml")

This step was not required in the Lissajous example because the server uses
standard heuristics to recognize common formats like PNG from the first 512
bytes of the response and generates the proper header. Allow the client to
specify values like height, width, and color as HTTP request parameters.
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth  = 600
	defaultHeight = 320
	defaultCells  = 100
	defaultColor  = "gray"
	xyrange       = 30.0
	angle         = math.Pi / 6
)

var (
	sin30, cos30 = math.Sin(angle), math.Cos(angle)
	clean        = flag.Bool("clean", false, "Filter non-finite polygons out")
)

func main() {
	http.HandleFunc("/", createSVG)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func createSVG(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	width := defaultWidth
	height := defaultHeight
	cells := defaultCells
	color := defaultColor
	for k, v := range r.URL.Query() {
		if len(v) == 0 || len(v) > 1 {
			log.Printf("Unexpected query param value for %s: %v", k, v)
			continue
		}

		value := v[0]
		var err error = nil
		switch k {
		case "width":
			width, err = strconv.Atoi(value)
		case "height":
			height, err = strconv.Atoi(value)
		case "cells":
			cells, err = strconv.Atoi(value)
		case "color":
			color = value
			err = nil
		default:
			log.Printf("Unexpected query param %s\n", k)
			err = nil
		}

		if err != nil {
			log.Printf("Error parsing query param %s with value %s: %g\n", k, value, err)
			continue
		}
	}
	fmt.Fprintf(w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: %s; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", color, width, height)

	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j, cells, width, height)
			bx, by := corner(i, j, cells, width, height)
			cx, cy := corner(i, j+1, cells, width, height)
			dx, dy := corner(i+1, j+1, cells, width, height)

			fmt.Fprintf(
				w,
				"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i int, j int, cells int, width int, height int) (float64, float64) {
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2D SVC Canvas (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x float64, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Cos(r) / r
}
