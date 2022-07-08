package main

import (
	"fmt"
)

// 求多个参数的和; args参数只能放到形参的最后, 和JS一样
// args的数据类型是一个切片
func add(args ...int) int {
	fmt.Printf("%T", args)
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	res := add(10, 20, 30, 40)
	fmt.Println(res)
}
