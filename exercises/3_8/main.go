/*
i
Exercise 3.8
Rendering fractals at high zoom levels demands great arithmetic precision.
Implement the same fractal using four different representations of numbers,
complex64, complex128, big.Float, big.Rat (the latter two types are from the
math/big package, Float uses arbitrary but bounded precision floating-point;
Rat uses unbounded precision rational numbers). How do they compare in
performance and memory usage? At what zoom-level do rendering artefacts become
visible?
*/
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/big"
	"math/cmplx"
	"math/rand"
	"os"

	"github.com/devscouse/go-programming-language/exercises/3_8/types"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, 2, 2
	minDistToRoot          = 1e-4
)

var (
	dtype      = flag.String("dtype", "complex128", "Accepted values are `complex64`, `complex128`, `BigFloat` and `BigRat`")
	size       = flag.Int("size", 1000, "The size of the render")
	precision  = flag.Uint("precision", 128, "The precision of the data type (if supported)")
	iterations = flag.Uint("iter", 100, "The number of iterations")
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

func newtonsMethodComplex64(z complex64, rootPoints [4]complex64, rootColors [4]color.RGBA) color.Color {
	for n := range *iterations {
		for i, rootPoint := range rootPoints {
			if cmplx.Abs(complex128(z-rootPoint)) < minDistToRoot {
				color := rootColors[i]
				color.A = 255 - uint8(255*n / *iterations)
				return color
			}
		}
		if cmplx.Abs(complex128(z)) < 1e-10 {
			z += complex(1e-10, 1e-10)
		}
		z3 := z * z * z
		z4 := z3 * z
		z = (3*z4 + 1) / (4 * z3)
	}
	return color.White
}

func newtonsMethodComplex128(z complex128, rootPoints [4]complex128, rootColors [4]color.RGBA) color.Color {
	for n := range *iterations {
		for i, rootPoint := range rootPoints {
			if cmplx.Abs(z-rootPoint) < minDistToRoot {
				color := rootColors[i]
				color.A = 255 - uint8(255*n / *iterations)
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

func newtonsMethodBigFloat(z types.BigFloatComplex, rootPoints [4]types.BigFloatComplex, rootColors [4]color.RGBA) color.Color {
	minDistFloat := big.NewFloat(minDistToRoot).SetPrec(z.Prec)
	minVal := big.NewFloat(1e-10).SetPrec(z.Prec)
	minValComplex := types.NewBigFloatComplex(1e-10, 1e-10, z.Prec)

	for n := range *iterations {
		for i, rootPoint := range rootPoints {
			if z.Sub(rootPoint).Abs().Cmp(minDistFloat) < 0 {
				color := rootColors[i]
				color.A = 255 - uint8(255*n / *iterations)
				return color
			}
		}
		if z.Abs().Cmp(minVal) < 0 {
			z = z.Add(minValComplex)
		}

		z2 := z.Mul(z)
		z3 := z2.Mul(z)
		z4 := z2.Mul(z2)

		numerator := z4.MulScalar(3.0).Add(types.NewBigFloatComplex(1, 0, z.Prec))
		denominator := z3.MulScalar(4.0)
		z = numerator.Div(denominator)
	}
	return color.White
}

func newtonsMethodBigRat(z types.BigRatComplex, rootPoints [4]types.BigRatComplex, rootColors [4]color.RGBA) color.Color {
	minDistRat := big.NewRat(1, 1e6)
	minValRat := big.NewRat(1, 1e10)
	minValComplex := types.NewBigRatComplex(1e-10, 0)

	const1 := types.NewBigRatComplex(1, 0)
	const3 := big.NewRat(3, 1)
	const4 := big.NewRat(4, 1)

	for n := range *iterations {
		log.Printf("Iteration %d started", n)
		for i, rootPoint := range rootPoints {
			if z.Sub(rootPoint).AbsSquared().Cmp(minDistRat) < 0 {
				color := rootColors[i]
				color.A = 255 - uint8(255*n / *iterations)
				return color
			}
		}
		if z.AbsSquared().Cmp(minValRat) < 0 {
			z = z.Add(minValComplex)
		}

		log.Printf("Iteration %d adjusting z value", n)

		z2 := z.Mul(z)
		z3 := z2.Mul(z)
		z4 := z2.Mul(z2)

		z = z4.MulScalar(const3).Add(const1).Div(z3.MulScalar(const4))

		// numerator := z4.MulScalar(const3).Add(types.NewBigRatComplex(1, 0))
		// denominator := z3.MulScalar(const4)
		// z = numerator.Div(denominator)
		log.Printf("Iteration %d complete", n)
	}
	return color.White
}

func mainComplex64() {
	rootPoints := [4]complex64{1 + 0i, -1 + 0i, 0 + 1i, 0 - 1i}
	rootColors := [4]color.RGBA{randomColor(), randomColor(), randomColor(), randomColor()}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))
	for py := range *size {
		for px := range *size {
			x, y := pixelPosToGlobalPos(float64(px), float64(py), float64(*size), float64(*size))
			z := complex(float32(x), float32(y))
			color := newtonsMethodComplex64(z, rootPoints, rootColors)
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}

func mainComplex128() {
	rootPoints := [4]complex128{1 + 0i, -1 + 0i, 0 + 1i, 0 - 1i}
	rootColors := [4]color.RGBA{randomColor(), randomColor(), randomColor(), randomColor()}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))
	for py := range *size {
		for px := range *size {
			x, y := pixelPosToGlobalPos(float64(px), float64(py), float64(*size), float64(*size))
			z := complex(x, y)
			color := newtonsMethodComplex128(z, rootPoints, rootColors)
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}

func mainComplexBigFloat() {
	rootPoints := [4]types.BigFloatComplex{
		types.NewBigFloatComplex(1, 0, *precision),
		types.NewBigFloatComplex(-1, 0, *precision),
		types.NewBigFloatComplex(0, 1, *precision),
		types.NewBigFloatComplex(0, -1, *precision),
	}
	rootColors := [4]color.RGBA{randomColor(), randomColor(), randomColor(), randomColor()}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))
	for py := range *size {
		for px := range *size {
			log.Printf("computation for pixel (%d, %d) complete", px, py)
			x, y := pixelPosToGlobalPos(float64(px), float64(py), float64(*size), float64(*size))
			z := types.NewBigFloatComplex(x, y, *precision)
			color := newtonsMethodBigFloat(z, rootPoints, rootColors)
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}

func mainComplexBigRat() {
	rootPoints := [4]types.BigRatComplex{
		types.NewBigRatComplex(1, 0),
		types.NewBigRatComplex(-1, 0),
		types.NewBigRatComplex(0, 1),
		types.NewBigRatComplex(0, -1),
	}
	rootColors := [4]color.RGBA{randomColor(), randomColor(), randomColor(), randomColor()}

	img := image.NewRGBA(image.Rect(0, 0, *size, *size))
	for py := range *size {
		for px := range *size {
			x, y := pixelPosToGlobalPos(float64(px), float64(py), float64(*size), float64(*size))
			z := types.NewBigRatComplex(x, y)
			color := newtonsMethodBigRat(z, rootPoints, rootColors)
			log.Printf("computation for pixel (%d, %d) complete", px, py)
			img.Set(px, py, color)
		}
	}

	png.Encode(os.Stdout, img)
}

func main() {
	flag.Parse()
	switch *dtype {
	case "complex64":
		mainComplex64()
	case "complex128":
		mainComplex128()
	case "BigFloat":
		mainComplexBigFloat()
	case "BigRat":
		mainComplexBigRat()
	default:
		os.Exit(1)
	}
}
