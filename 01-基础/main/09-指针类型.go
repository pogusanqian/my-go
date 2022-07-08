package main

import (
	"fmt"
)

/**
* 1. 指针表示存储内存地址的变量
* 2. &num: 获取num变量的指针, &符号的优先级要低于.
* 3: *ptr: 获取ptr指针指向内存地址的值, 也就是变量num的值
* 4. *int: int类型的指针
* 5. 值类型: int, float, boolean, string, 数组, 结构体
* 6. 引用类型: 指针, 切片, 管道, map, interface都是引用数据类型
*/
func main() {
	// 创建了一个num变量, 系统系统会在内存中划分一块空间给num变量, 这块空间有一个编号, 这个编号就是内存地址
	var num int = 100

	// &num获取num变量的内存地址, &表示取地址符号
	// *int: 声明了一个int类型的指针变量; 可以得出指针其实就是一个存储内存地址的变量
	var ptr *int = &num

	// *ptr: 获取内存地址的值
	fmt.Println(ptr, *ptr) // 0xc000014098, 100

	*ptr = 1000
	// 通过指针修改内存空间中的值, 发现基本类型的值会被修改(调用方法传参A时, 可以传递指针类型, 这样就相当于传递了一个引用类型, 就可以在方法中直接修改A的值)
	fmt.Println(*ptr, num) // 1000, 1000

	// 指针变量也可以有指针, &ptr就表示指针变量的指针
	fmt.Println(ptr, &ptr, *&ptr) // 0xc000014098 0xc000006028 0xc000014098
}
