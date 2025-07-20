// Package unitconv provides types and methods for converting values between common units
// of measurement.
package unitconv

import "fmt"

type (
	Celsius    float64
	Fahrenheit float64
	Feet       float64
	Meters     float64
	Pounds     float64
	Kilograms  float64
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string     { return fmt.Sprintf("%gm", m) }
func (p Pounds) String() string     { return fmt.Sprintf("%glb", p) }
func (k Kilograms) String() string  { return fmt.Sprintf("%gkg", k) }

func (c Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func (f Fahrenheit) ToCelsius() Celsius    { return Celsius((f - 32) * 5 / 9) }
func (f Feet) ToMeters() Meters            { return Meters(f * 0.3048) }
func (m Meters) ToFeet() Feet              { return Feet(m * 3.28084) }
func (p Pounds) ToKilograms() Kilograms    { return Kilograms(p * 0.4535924) }
func (k Kilograms) ToPounds() Pounds       { return Pounds(k * 2.204623) }
