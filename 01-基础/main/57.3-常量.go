package main

import (
	"fmt"
)

/**
* 1. 常量只能是布尔, 数字, 字符串三种类型
* 2. go语言中的常量是不需要大写的
 */
func main() {
	const name string = "张三"

	// iota表示0, b, c依次加一
	const (
		a = iota
		b
		c
	)

	fmt.Println(name, a, b, c)
}
