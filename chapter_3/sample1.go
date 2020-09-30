package main

import "fmt"

func main() {

	//变量三要素， 类型，名，值

	var a int
	//声明后不赋值，使用默认值，int/float-->0, string--->"", bool--->false, 引用类型--->nil
	fmt.Println(a)

	//先声明后赋值
	var i int
	i = 10
	fmt.Println("i=", i)

	//类型推断
	var b = 10

	//省略var，:= 声明并赋值+类型推断
	c := 10

	//同时声明多个变量
	var cc, aa, bb int
	n1, n2, n3 := 10, 1, 1
	n4, n5, n6 := false, "false", 0

	var (
		r = 1
		e = 2
		x = "x"
	)

	fmt.Println(a,i,b,c,cc,aa,bb,n1,n2,n3,n4,n5,n6,x,e,r)
}
