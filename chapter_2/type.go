package main

import "fmt"

func main() {

	var i int32 = 100
	var n1  float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v n1=%v n2=%v n3=%v", i, n1, n2, n3)
	//Go中数据可以从范围大--->范围小的，也可以：范围小的-->范围大的
	//转换的只是数值，数据本身并没有发生变化
	//由大范围--->小范围，按照溢出处理


}
