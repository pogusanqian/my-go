package main

import (
	"fmt"
)

/**
* go中是没有while和dowhile的语法的, 可以使用for循环来实现
 */
func main() {
	var flag = 1
	var flag2 = 10

	// 使用for实现while循环(先执行判断, 再执行循环语句)
	for {
		if flag > 10 {
			break // break可以跳出循环, 不用使用return
		}
		fmt.Println("好好学习, 天天向上")
		flag++
	}

	// 使用for实现do...while循环(限制性循环体, 再执行判断)
	for {
		fmt.Println("Hello World")
		flag2++
		if flag2 > 10 {
			break
		}
	}
}
