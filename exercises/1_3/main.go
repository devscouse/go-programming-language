// Exercise 1.3: Experiment to measure the difference in running time between
// our potentially inefficient versions and the one that uses strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	echo1()
	fmt.Printf("\necho1 took %d microseconds\n", time.Since(start).Microseconds())

	start = time.Now()
	echo2()
	fmt.Printf("\necho2 took %d microseconds\n", time.Since(start).Microseconds())

	start = time.Now()
	echo3()
	fmt.Printf("\necho3 took %d microseconds\n", time.Since(start).Microseconds())
}
