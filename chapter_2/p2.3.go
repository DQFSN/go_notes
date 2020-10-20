package main

import (
	"go_notes/chapter_2/popcount"
	"fmt"
	"time"
)

func main() {
	begin := time.Now().UnixNano()

	popcount.PopCount(9999)

	end1 := time.Now().UnixNano()

	popcount.PopCount2_3(9999)

	end2 := time.Now().UnixNano()

	fmt.Printf("first: %d\n", end1 - begin)
	fmt.Printf("second: %d\n", end2 - end1)
}
