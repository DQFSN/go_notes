package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)


func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Get %s", err)
			os.Exit(1)
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:read %s : %s", url, err)
			os.Exit(1)
		}

		fmt.Println("read :", n)
	}
}