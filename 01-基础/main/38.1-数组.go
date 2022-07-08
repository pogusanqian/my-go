package main

import (
	"fmt"
)

/**
* 1. 数组在Go中传参的时候, 是值传递; JS中的是引用传递
* 2. 数组的地址就是第一个元素的地址
* 3. 数组是相同数据类型的组合, 一旦声明之后, 数组的长度是不能动态变化的(JS中的数组可以存储不同的数据类型)
* 4. 数组传递时是值传递, 不会影响原本的值(JS是引用传递)
 */
func main() {
	// 声明一个数组, 长度是5, 元素类型是int, 每个元素的默认值是0
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	fmt.Println(arr1, arr1[0])
	// 输出数组的地址(使用println是无法输出的, 数组元素的内存空间是连续的)
	fmt.Printf("%p, %p, %p", &arr1, &arr1[0], &arr1[1])

	// 初始化方式
	var arr2 [3]int = [3]int{1, 2, 3} // 这里声明的数组长度是不能超过3的, 否则会发生数组越界
	var arr3 = [3]int{11, 2, 3}
	var arr4 = [...]int{111, 2, 3, 4, 5}
	var arr5 = [3]string{2: "王五", 0: "张三", 1: "李四"}

	fmt.Printf("%v\n %v\n %v\n %v\n", arr2, arr3, arr4, arr5)

}
