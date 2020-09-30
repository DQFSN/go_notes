package main

import "fmt"


func main() {
	
	var c1 byte = 'c'
	var c2 byte = '0'

	//默认输出码值
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	//输出字符
	fmt.Printf("c1 = %c, c2 = %c\n", c1, c2)

	var c3 int = '背'
	fmt.Printf("c3 = %c, 对应码值=%d\n", c3, c3)
	var cc = 22269
	fmt.Printf("22269 对应的字符是%c\n",cc)


}
