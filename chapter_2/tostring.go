package main

import "fmt"

func main() {
	var num int = 99
	var num2 float32 = 1.2
	var b bool = false
	var mychar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num)
	fmt.Println(str,num2)

	str = fmt.Sprintf("%t", b)
	fmt.Println(str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Println(str)
}
