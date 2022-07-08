package main

import (
	"fmt"
)

/**
* 1. 切片是一个引用类型, 包含三个元素, 一个指向一个数组的指针, 一个是切片长度, 一个是切片容量
* 2. 切片的长度是可以动态变化的
* 3. 切片本身就是引用数据类型, 所以根据&符号获取不到指针
 */
func main() {
	// 声明一个切片, 然后使切片指向一个数组[1:3]表示是下标[1,3)
	var myslice []int // // 注意这里不是声明了一个数组, 而是声明了一个切片, 数组是需要执行长度的
	var intArr = [5]int{0, 1, 2, 4, 5}
	myslice = intArr[1:3]

	intArr[2] = 22 // 修改数组的值, 也是修改了切片的地址
	fmt.Println("切片值", myslice)
	fmt.Println("切片长度", len(myslice))
	fmt.Println("切片容量", cap(myslice))

	// 使用make创建一个容量是10, 长度是5的切片; 通过make创建的切片, 切片指向的数组程序员是不可见的, 使用go引擎来维护
	var myslice2 = make([]int, 5, 10)
	fmt.Println(myslice2)

	// 定义切片并指定具体的数组
	var myslice3 []string = []string{"张三", "李四", "王五"}
	fmt.Println(myslice3)
}
