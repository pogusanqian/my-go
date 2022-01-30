package main

import (
	"fmt"
)

func main() {
	// myslice2是由于myslice再切片出来的, 他们指向同一个内存空间
	myslice := []string{"张三", "李四", "王五"}
	myslice2 := myslice[:]

	myslice2[0] = "zhangsan"
	fmt.Println(myslice, myslice2)

}
