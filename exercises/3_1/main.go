// Surface computes an SVG rendering of a 3D surface function
// Exercise 3.1: If the function f returns a non-finite float64 value, the SVG
// file will contain invalid <polygon> elements (although many SVG renderers
// handle this gracefully). Modify the program to skip invalid polygons.
package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	width   = 600
	height  = 320
	cells   = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale  = height * 0.4
	angle   = math.Pi / 6
)

var (
	sin30, cos30 = math.Sin(angle), math.Cos(angle)
	clean        = flag.Bool("clean", false, "Filter non-finite polygons out")
)

func main() {
	flag.Parse()
	fmt.Printf(
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if *clean {
				arr := []float64{ax, ay, bx, by, cx, cy, dx, dy}
				ok := true

				for _, val := range arr {
					if math.IsInf(val, 0) || math.IsNaN(val) {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}
			}

			fmt.Printf(
				"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

// Find surface point (sx, sy) at corner of cell (i, j)
func corner(i int, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2D SVC Canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x float64, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Cos(r) / r
}
