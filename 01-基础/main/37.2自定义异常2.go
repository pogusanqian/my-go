package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误..")
	}
}

// TODO 使用panic捕捉异常有什么用, 程序依然会终止的
func test() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test()继续执行....")
}

func main() {
	test()
	fmt.Println("main()下面的代码...")
}
