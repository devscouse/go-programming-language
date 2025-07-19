// Excercise 1.8: Modify fetch to add the prefix http:// to each argument URL
// if it is missing. You might want to use strings.HasPrefix.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		nbytes, err := io.Copy(os.Stdout, res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("fetch %s returned %d bytes\n", url, nbytes)
	}
}
