package util

import (
	"fmt"
)

type People2 struct {
	Name string
	Age  int
}

type Student2 struct {
	Name string
	Age  int
}

type MiddleStudent2 struct {
	People2
	Student2
	Age int
}

func (stu *Student2) Show() {
	fmt.Println(stu.Name, stu.Age)
}

// 如果有两个重名的方法, 在在会报错ambiguous selector stu.Name
func (stu *MiddleStudent2) Show() {
	fmt.Println(stu.Age)
}
