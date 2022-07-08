package main

import (
	"fmt"
)

// 数组传递时将原数组的值copy到了myarr, 修改myarr是不会影响arr的;
// 如果想要在test方法中影响arr的值, 可以传一个指针数组过来
// 这里传递过来数组长度只能是3
func test(myarr [3]string) {
	myarr[0] = "zhangsan"
	myarr[1] = "lisi"
	myarr[2] = "wangwu"
}

// 传递指针数组, 修改值
func test2(myarr *[3]string) {
	(*myarr)[0] = "zhangsan"
	(*myarr)[1] = "lisi"
	(*myarr)[2] = "wangwu"
}

func main() {
	var arr = [3]string{"张三", "李四", "王五"}

	test2(&arr)
	fmt.Println(arr)
}
