package main

import (
	"fmt"
)

func calc(num1 int) {
	num1 = num1 + 100
	fmt.Println(num1) // 200
}

func main() {
	var num = 100
	calc(num)        // 调用calc之后, 是不会影响原本的num值的; 这种传参相当于把num的值赋值给了calc函数中的num1变量
	fmt.Println(num) // 100
}
