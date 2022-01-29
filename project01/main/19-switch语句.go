package main

import (
	"fmt"
)

/**
* 1. switch语句中的每一个case不需要使用break结束, go自动提供了
* 2. switch/case后面可以是一个表达式, 甚至是一个由返回值的函数(类似于js), 但是又一点需要注意, 就是case中的值必须要和变量的数据类型一样
 */
func main() {
	var num byte
	// var num2 int32 = 10
	fmt.Println("请输入一个数字:")
	fmt.Scanln(&num)

	// 变量卸载swithc后面
	switch num {
	// 这里的num2和num的数据类型不一样, 再比较的时候是会报错的, 我们可以直接写成20
	// case num2:
	// 	fmt.Println("我是10")
	case 20:
		fmt.Println("我是20")
	case 30:
		fmt.Println("我是30")
	case 40, 50:
		fmt.Println("我是40, 50")
	default:
		fmt.Println("我是默认值")
	}

	// 变量写在case后面
	switch {
	case num < 30:
		fmt.Println("我小于30")
	case num < 50:
		fmt.Println("我小于50")
	default:
		fmt.Println("我大于50")
	}

	// switch的主动穿透
	switch {
	case num < 30:
		fmt.Println("我小于30")
		fallthrough
	case num < 50:
		fmt.Println("我小于50")
	default:
		fmt.Println("我大于50")
	}
}
