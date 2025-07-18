// Find a web site that produces a large amount of data. Investigate caching
// by running fetchall twcice in succession to see whether the reported time
// changes much. Do you get the same content each time? Modify fetchall to print
// its output to a file so it can be examined.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fetchall()
	fetchall()
}

func fetchall() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // recieve from ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	filename := start.String()
	f, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("while creating %s: %v", filename, err)
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
