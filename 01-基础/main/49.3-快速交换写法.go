package main

import (
	"fmt"
)

func main() {
	var num1 = 100
	var num2 = 200
	// 快速交换num1和num2的值
	num1, num2 = num2, num1
	fmt.Println(num1, num2)
}
