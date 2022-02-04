package main

import (
	"fmt"
)

// defer最主要的价值就是当函数执行完毕后, 可以即使的释放资源
func sum(num1 int, num2 int) int {
	// defer会将语句压入到独立的栈中, 等函数执行完毕后, 在从defer栈中按照先进后出的原则取出语句进行执行
	// dfer压入的变量时, 会将变量的值copy进去
	defer fmt.Println("===========defer1", num1) // 这里输出的是10
	defer fmt.Println("===========defer2", num2) // 这里输出的是20

	num1 = num1 + 100
	num2 = num2 + 100
	fmt.Println("===========sum函数执行", num1, num2) // 110, 120
	return num1 + num2
}

func main() {
	res := sum(10, 20)
	fmt.Println("============", res) // 最后执行
}
