package main

import (
	"fmt"
	"reflect"
)

/**
* Go语言的主要应用场景就是大规模的后端服务程序, 网络通信是服务器之间必不可少的一部分
* 1. TCP socket编程: 网络编程的主流, 之所以叫socket是因为底层使用的是tcp/ip协议, 如qq聊天
* 2. http编程: 主要用于B/S架构程序, http编程底层依然使用的是tcp socket编程如网页版的京东商城
 */
func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}

	TestStruct(a)
}
