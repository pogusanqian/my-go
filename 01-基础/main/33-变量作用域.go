package main

import (
	"fmt"
)

/**
* 1. {}声明的变量叫做局部变量, 只能在{}中声明; 和JS中的块作用域一样
* 2. 函数外部声明的变量叫做全局变量, 在整个包中都有效, 如果首字母大小则在整个程序中都有效
 */
func main() {
	if true {
		var name = "张三"
	}
	fmt.Println(name)
}
