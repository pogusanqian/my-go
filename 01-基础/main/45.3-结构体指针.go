package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   byte
	Color string
}

// 结构体的指针值并不能直接打印, 需要使用%p参数
// 结构体第一个属性值的指针其实就是结构体的指针
// 结构体属性的指针值是连续的
func main() {
	var cat = Cat{"小黑", 3, "黑色"}
	var ptr *Cat = &cat

	fmt.Printf("%T\n", ptr)
	fmt.Printf("%p %v\n", ptr, ptr) // 0xc000114480 &{小黑 3 黑色}
	fmt.Printf("%p %p %p %p\n", ptr, &cat.Name, &cat.Age, &cat.Color) // 0xc000114480 0xc000114480 0xc000114490 0xc000114498
}
