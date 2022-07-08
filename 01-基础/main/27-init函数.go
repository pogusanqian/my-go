package main

import "fmt"

// init函数通再import导入的时候就会进行加载
func init() {
	fmt.Println("============init", name)
}

func main() {
	fmt.Println("=============main", name)
}

// 全局变量会在init函数之前调用
var name = "张三"
