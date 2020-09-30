package main

import "fmt"

func main() {
	var address string = "北京长城 110 "
	fmt.Println(address)
	//字符串一旦赋值，无法修改，address[0] = 'd',错误
	//字符串的两种表示形式
	//1、双引号，会识别转义字符
	//2、反引号，以字符串的原生形式输出，不转义字符，可以防止攻击，输出源代码等

	str := `package main
	import "fmt"

	func main() {
	}
	`

	str1 := "hellor" + "world" + 
	" !"

	fmt.Println(str)
	fmt.Println(str1)

}
