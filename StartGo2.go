package main

import (
	"fmt"
)

var a bool

var c int = 10

//Go的基本类型 bool string int int8 int32 int64
//uint uint8 uint32 uint64
// byte  uint8的别名
// rune int32的别名
//float32 float64
//complex64 complex128

// 类型转换   隐式类型转换比较复杂
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

const World = "我想要摆烂的自由"

//常量的声明与变量类似，只不过是使用const关键字，不能用 := 定义
//常量可以是字符 字符串 布尔值 或 数值
//数值常量是高精度的值 并且 一个未指定类型的常量由上下文来决定其类型

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {

	var b int
	//没有明确初始值的变量声明会被赋予它们的 零值
	//数值类型为0 布尔类型为false 字符串为“”
	fmt.Println(a, b)

	var d, e bool = true, true
	fmt.Println(c, d, e)

	k := 6 //简洁赋值语句：=可在类型明确的地方 代替var声明
	// 函数外的每个语句都必须以关键字开始（var func等等） 因此 ：= 结构不能在函数外使用
	fmt.Println(k)

	fmt.Println(u)

	i := 42 //不能在函数外使用
	f := float64(i)
	u := uint(f)
	fmt.Println(u)

	fmt.Println(World)
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needFloat(y float64) float64 {
	return y * 0.1
}
