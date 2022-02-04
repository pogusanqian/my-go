package main

import (
	"fmt"
)

// 返回一个值
func numAdd(num1 float64, num2 float64) float64 {
	return num1 + num2
}

// 返回两个值
func calc(num1 float64, num2 float64) (float64, float64, float64, float64) {
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	return sum, sub, mul, div
}

func main() {
	fmt.Println(numAdd(20, 30))

	// _只是一个占位符, 忽略返回值, 只能占一个返回值
	res1, res2, _, _ := calc(20, 30)
	fmt.Println(res1, res2)
}
