// Exercise 3.2: Experiment with visualisations of other functions from the
// math package. Can you produce an egg box, moguls, or a saddle.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	width   = 600
	height  = 600
	cells   = 100
	xyrange = 1.0
	xyscale = width / 2 / xyrange
	zscale  = height * 0.7
	angle   = math.Pi / 6
)

var (
	sin30    = math.Sin(angle)
	cos30    = math.Cos(angle)
	plotFunc = flag.String("function", "saddle", "The function you want to plot")
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
	if *plotFunc == "saddle" {
		return saddle(x, y)
	}
	if *plotFunc == "eggbox" {
		return eggbox(x, y)
	}
	fmt.Fprintf(os.Stderr, "Unsupported plotFunc %s\n", *plotFunc)
	os.Exit(1)
	return 0
}

func eggbox(x float64, y float64) float64 {
	return 2 * math.Sin(x*1.5) * math.Sin(y*1.5)
}

func saddle(x float64, y float64) float64 {
	return x*x + y*y*y
}
