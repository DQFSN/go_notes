package main

import "go_notes/chapter_4/countdiffbit"
import "fmt"

func main() {
	fmt.Println(countdiffbit.CountDiffBit("x","x"))
	array := &[...]int{1,2,3,4,5,6}
	countdiffbit.Reverse(array)
	fmt.Println(*array)
}


