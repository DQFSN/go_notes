package main

import "fmt"
import "unsafe"

func main() {
	//一个字节，-128～127
	var i int8 = 127
	fmt.Println(i,i+1,i+2)
	var j uint8 = 255
	fmt.Println(j,j+1)
	//rune 和 int32 等价，byte和uint8等价
	//int 和uint的大小与系统有关，64位系统是8字节，32位是4字节
	var n1 = 100
	fmt.Printf("n1的类型%T, n1的字节大小%d\n", n1, unsafe.Sizeof(n1))
}

