// Exercise 1.2: Modify the echo program to print the index and value of
// each of its arguments, one per line
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i) + " " + arg)
	}
}
