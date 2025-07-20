// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingC = 100.0

func main() {
	c := boilingC
	f := (c * 9 / 5) + 32
	fmt.Printf("boiling point = %g degF or %g degC\n", f, c)
}
