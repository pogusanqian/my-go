package main

import (
	"fmt"
)

func main() {
	// 除法运算
	fmt.Println(10 / 4)   // 2 去掉了浮点数部分
	fmt.Println(10.0 / 4) // 2.5 如果想要保留小数部分, 需要要浮点数参与运算
	fmt.Println(10 / 4.0) // 2.5 如果想要保留小数部分, 需要要浮点数参与运算

	// 取模运算: a % b = a - (a / b) * b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  // -1
	fmt.Println(10 % -3)  // 1
	fmt.Println(-10 % -3) // -1

	// ++ --只能单独使用, 不能进行赋值
	// var num = 3++					// 报错
	// fmt.Println(num)
}
