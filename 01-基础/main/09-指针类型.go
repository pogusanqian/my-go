package main

import (
	"fmt"
)

/**
1. 指针指向变量的真正内存地址
2. &num获取变量的指针
3. *int表示一个int类型的指针
4. 值类型: int, float, boolean, string, 数组, 结构体都有对应的指针
	  值类型变量直接存储值, 内存通常在栈中分配
5. 引用类型: 指针, 切片, 管道, map, interface都是引用数据类型
    引用类型指向一个地址, 这个地址对应的空间才是真正的值
*/
func main() {
	var num = 100

	// 声明指针变量, 使用*; 获取指针是使用的是&num
	var ptr *int = &num
	fmt.Printf("%v\n", ptr)

	// 通过指针修改内存空间中的值, 发现基本类型的值会被修改
	*ptr = 1000
	fmt.Println(num)

	// 使用num引用修改基本类型的值, 和通过指针修改是一样的
	num = -222
	fmt.Println(num, *ptr)

	// 字符类型指针
	var c byte = 'a'
	fmt.Println(&c)
}
