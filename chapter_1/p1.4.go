package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(filename,err)
		}

		input := bufio.NewScanner(file)

		for input.Scan() {
			counts[input.Text()]++
			if counts[input.Text()] > 1 {
				fmt.Println(input.Text(), filename)
			}
		}
	}
}
