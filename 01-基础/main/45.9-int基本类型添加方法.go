package main

import (
	"fmt"
)

type Myint int

func (i *Myint) addTen() {
	// 注意这里的必须写成*i, 因为i本身时指针类型, 指针类型是不能进行加法运算的
	*i = *i + 10
}

func main() {
	var i Myint = 1
	i.addTen()
	fmt.Println(i)
}
