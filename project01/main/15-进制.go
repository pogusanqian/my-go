package main

import (
	"fmt"
)

func main() {
	var num1 = 24   // 十进制表示数据
	var num2 = 011  // 8进制表示数据
	var num3 = 0x11 // 16进制表示数据

	fmt.Printf("%b \n", num1) // 二进制输出, go语言中不能直接表示二进制, 但是可以通过函数转换
	fmt.Println(num2)
	fmt.Println(num3)

}
