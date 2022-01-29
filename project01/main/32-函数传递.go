package main

import (
	"fmt"
)

/**
* 函数参数传递种类
*   值传递: 将值copy之后, 传递给了方法的形参
* 	引用传递: 将引用地址copy传递了过去; 引用传递效率比较高, 因为只是一个引用地址
*
* Go语言比较特殊, 基本类型以及数组和结构体被当成了值传递了过去, JS中的数组时传递过去的引用
* 		引用类型传递: 指针, slice切片, map, 管道chan, 接口interface
 */
func main() {
	fmt.Println("========")
}
