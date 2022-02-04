package main

import (
	"fmt"
)

/**
* 使用append扩容之后, 会返回一个新的切片, 我们会发现&myslice[0]不是同一个值
 */
func main() {
	myslice := []string{"张三", "李四", "王五"}
	fmt.Println(&myslice, &myslice[0])

	// 给切片添加元素进行扩容
	myslice = append(myslice, "赵六", "李七")
	fmt.Println(&myslice, &myslice[0])

	// 给切片添加切片扩容
	myslice = append(myslice, myslice...)
	fmt.Println(&myslice, &myslice[0])
}
