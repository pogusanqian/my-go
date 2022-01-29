package main

import "fmt"

// init 会在main函数之前进行调用, 已完成初始化公共
func init() {
	fmt.Println("============init", name)
}

func main() {
	fmt.Println("=============main", name)
}

// 全局变量会在init函数之前调用
var name = "张三"
