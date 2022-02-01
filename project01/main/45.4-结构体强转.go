package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

type People struct {
	name string
	age  int
}

/**
* 1. 结构体强转的时候, 要求两个结构体的字段名, 字段个数, 字段类型需要完全一致才能就行强转
* 2. 使用type会对结构体从重新定义, 是两个不同的数据类型, 如 type Myint int 和int是不同的数据类型
 */
func main() {
	var stu = Student{"张三", 23}
	var peo = People(stu) // 将Student强转成了People类型
	fmt.Println(peo)
}
