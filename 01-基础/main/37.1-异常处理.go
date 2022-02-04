package main

import (
	"fmt"
)

/**
* 1. go语言中是没有try...catch的概念的, 使用的是使用defer + recover 来捕获和处理异常
 */
func test() {
	defer func() {
		err := recover() // recover可以捕捉异常
		if err != nil {
			fmt.Println("抓到异常:", err)
		}
	}()

	// TODO 为什么直接使用10 / 0的异常根本抓不到
	// fmt.Println(10 / 0)

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	fmt.Println("===============")
}
