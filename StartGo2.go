package main

import "fmt"

var a bool

var c int = 10

//Go的基本类型 bool string int int8 int32 int64
//uint uint8 uint32 uint64
// byte  uint8的别名
// rune int32的别名
//float32 float64
//complex64 complex128

func main() {

	var b int
	fmt.Println(a, b)

	var d, e bool = true, true
	fmt.Println(c, d, e)

	k := 6 //简洁赋值语句：=可在类型明确的地方 代替var声明
	// 函数外的每个语句都必须以关键字开始（var func等等） 因此 ：= 结构不能在函数外使用
	fmt.Println(k)
}
