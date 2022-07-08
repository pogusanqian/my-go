package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// 使用指针给结构体绑定方法
func (stu *Student) show() {
	stu.Name = "zhangsan"
}

func main() {
	var stu = Student{"张三", 23}
	stu.show() // 这里其实也是使用的是指针进行的调用, 只不过被简化了; &stu.show()
	fmt.Println("--------", stu.Name)
}
