package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	begin := time.Now().UnixNano()

	for i := 0; i < 100000; i++ {
		strings.Join(os.Args[1:], " ")
	}
	end := time.Now().UnixNano()

	for i := 0; i < 100000; i++ {
		str := ""
		for _, v := range os.Args[:] {
			str = v + " "
		}
		_ = str
	}


	end2 := time.Now().UnixNano()

	fmt.Println("join:",(end - begin) / 100000.0, " string + ：", (end2 - end) / 100000.0)
}
