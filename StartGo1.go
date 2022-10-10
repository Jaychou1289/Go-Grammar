package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func reduce(x, y int) int { //当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
	return x - y
}

func swap(x, y string) (string, string) { //函数可以返回任意数量的返回值
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //没有参数的retrun语句 返回已命名的返回值。
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(math.Sqrt(7))

	fmt.Println(math.Pi) //如果一个名字是大写字母开头，则它是已导出的，否则是未导出的。
	//在导入一个包时，你只能引用其中已导出的名字

	fmt.Println(add(3, 7))

	fmt.Println(reduce(7, 2))

	fmt.Println(swap("枫", "陈"))

	fmt.Println(split(17))
}
