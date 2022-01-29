package main

import (
	"fmt"
)

// 这个函数返回的是一个函数
func addUpper() func(int) int {
	var n = 10
	return func(num int) int {
		n = n + num
		return n
	}
}

func main() {
	// 这里获取的一个函数, 函数中的n变量一直没有释放
	myFun := addUpper()

	fmt.Println(myFun(1)) // 11
	fmt.Println(myFun(2)) // 13
	fmt.Println(myFun(3)) // 16
}
