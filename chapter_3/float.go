package main

import "fmt"

func main() {
	var price float32 = 1.223
	fmt.Println("price = ", price)
	//float32--->单精度，4字节， float64--->双精度 8字节
	//浮点数都是有符号的
	//尾数部分有可能丢失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 = ", num1, "num2 = ", num2)
	//所以我们一般用float64

	num3 := 1.21
	num4 := .121
	num5 := 1.2e2
	num6 := 1.2e-2

	fmt.Println(num3, num4, num5, num6)

}
