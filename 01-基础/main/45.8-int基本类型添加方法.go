package main

import (
	"fmt"
)

type Myint int

func (i *Myint) add() {
	// TODO 为什么这里的*号不能简写
	*i = *i + 100
}

func main() {
	var i Myint = 1
	(&i).add()
	fmt.Println(i)
}
