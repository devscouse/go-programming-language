package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/devscouse/go-programming-language/exercises/2_2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		ft := unitconv.Feet(t)
		m := unitconv.Meters(t)
		lb := unitconv.Pounds(t)
		kg := unitconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n", f, f.ToCelsius(), c, c.ToFahrenheit())
		fmt.Printf("%s = %s, %s = %s\n", ft, ft.ToMeters(), m, m.ToFeet())
		fmt.Printf("%s = %s, %s = %s\n", lb, lb.ToKilograms(), kg, kg.ToPounds())
	}
}
