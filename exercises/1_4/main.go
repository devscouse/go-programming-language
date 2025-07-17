// Excercise 1.4: Modify dup2 to print the names of all files in which each
// duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineCounts := make(map[string]int)
	fileLocations := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", lineCounts, fileLocations)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, lineCounts, fileLocations)
			f.Close()
		}
	}
	for line, n := range lineCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("\tFound in %s\n", fileLocations[line])
		}
	}
}

func countLines(f *os.File, fname string, counts map[string]int, filemap map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		filemap[line] += fname + ", "
	}
}
