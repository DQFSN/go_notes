package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)


func main() {

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Get %s", err)
			os.Exit(1)
		}

		fmt.Println("read :", resp.Status)
	}
}