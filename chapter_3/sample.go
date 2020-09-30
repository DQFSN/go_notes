package main

import "fmt"

func getVal(num1, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main() {
	
	sum, sub := getVal(3, 2)
	fmt.Println("sum=", sum, "sub=", sub)
	sum, _ = getVal(10, 2)
	fmt.Println("sum=", sum)
}


