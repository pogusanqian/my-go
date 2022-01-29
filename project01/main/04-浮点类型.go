package main

import "fmt"

/**
1. 浮点数都是由符号的
2. 浮点数可能会有精度丢失
3. 浮点数的默认数据类型是float64, 默认值也是0
*/
func main() {
	var num1 float32 = 1.00000001
	var num2 float64 = 1.00000001

	fmt.Println(num1) // 1 精度丢失
	fmt.Println(num2)

}
