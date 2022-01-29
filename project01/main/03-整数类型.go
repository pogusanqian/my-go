package main

import(
	"fmt"
	"unsafe"
) 

/**
	整数数据类型
		1. int: 有符号, 32位系统中占4个字节, 64位系统中占用8个字节, 所以int和int32是不同的数据类型
		2. uint: 无符号的int类型
		3. rune: 等价于int32
		4. byte: 1个字节空间, 存储数据类型时, 相当于时int8
		5. go语言中, var num = 100; 其中num的数据类型默认就是int数据类型
		6. 声明变量原则: 再程序正常行运行时, 尽量采用占用空间小的数据类型
*/
func main() {
	// 查看变量的数据类型
	fmt.Printf("%T", 100)
	fmt.Println()

	// 查看变量的内存大小
	fmt.Printf("%d", unsafe.Sizeof(100))
	fmt.Println()

	// 使用byte声明变量
	var num byte = 30
	fmt.Println(num)
	
}