package main

import "fmt"

func test() bool {
	fmt.Println("天天向上")
	return true
}

func main() {
	// 省略if的括号
	if 2 < 9 && test() {
		fmt.Println("好好学习")
	}
}
