package main

import (
	"fmt"
)

func main() {
	// 直接运行匿名函数
	var res = func(num1 int, num2 int) int {
		return num1 + num2
	}(20, 30)
	fmt.Println(res)

	// 将匿名函数赋值到变量上
	var myFun = func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println(myFun(30, 50))
}
