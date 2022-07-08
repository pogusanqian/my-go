package main

import (
	"fmt"
)

/**
* 1. go语言中是没有try...catch的概念的, 使用的是使用defer + recover来捕获, 使用panic抛出异常
* 2. 不需要再每一个函数中都defer异常, 直接再根目录或者其他父级目录recover异常, 这样也能捕捉到子函数的异常
* 	 注意此时子函数的异常, 会影响父函数的正常执行, 因为是在父函数中defer + recover捕捉了异常(相当于把try移动到了父函数, 扩大了范围, 健壮性就会弱一点)
*/
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("抓到异常:", err)
		}
	}()
	// 这里不要直接使用10/0, 因为这样的话, 编译阶段就会报错
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	fmt.Println("===============")
}
