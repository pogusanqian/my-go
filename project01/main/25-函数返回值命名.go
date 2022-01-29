package main

import (
	"fmt"
)

// 在返回值中声明sum和sub两个变量, 用来接收返回值; return不能忘记书写
func calc(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	res1, res2 := calc(30, 10)
	fmt.Println(res1, res2)
}
