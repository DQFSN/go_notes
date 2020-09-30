package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) < 1 {
		fmt.Println("no input file")
		return
	}else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f,counts)
			f.Close()

		}
	}

	for line,n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}

}