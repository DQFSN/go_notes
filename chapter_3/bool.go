package main

import "fmt"
import "unsafe"

func main() {

	var b = false
	fmt.Println(b)

	fmt.Println("bool的大小：", unsafe.Sizeof(b))
}
