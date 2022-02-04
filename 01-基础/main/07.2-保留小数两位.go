package main

import (
	"fmt"
)

func main() {
	num1 := 10 / 3.0
	fmt.Println(num1)          // 3.3333333333333335 精度丢失
	fmt.Printf("%.2f\n", num1) // 保留两位小数输出

	num2 := fmt.Sprintf("%.2f", 10/3.0) // 计算两位小数
	fmt.Println(num2)
}
