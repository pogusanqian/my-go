package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("抓到异常:", err)
		}
	}()

	// 如果抛出的异常不进行捕捉的话, 就会直接终止整个程序
	panic(errors.New("读取文件错误.."))
}

func main() {
	test()
	fmt.Println("main()下面的代码...")
}
