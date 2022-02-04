package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// 通过指针绑定时, stu就不在时值copy了
func (stu *Student) show() {
	// (*stu).Name = "张三"的简写
	stu.Name = "zhangsan"
}

func (stu Student) show2() {
	stu.Name = "李四"
}

func main() {
	var stu = Student{"张三", 23}
	stu.show() // 这里其实也是使用的是指针进行的调用, 只不过被简化了
	fmt.Println("--------", stu.Name)

	(&stu).show2() // 这里是进行的是值copy, 因为show2的方法是按照结构体进行的接受
	fmt.Println("--------", stu.Name)
}
