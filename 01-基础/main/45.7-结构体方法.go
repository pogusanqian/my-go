package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// 给结构体添加方法, stu只是一个副本, 如果在show方法中修改stu的name值, 是不会影响原本的值; 所以通常绑定方法的时候, 绑定的是指针
// show中如果有参数的话, 也是进行了值或引用copy
func (stu Student) show() {
	fmt.Println("=========", stu.Name)
	stu.Name = "zhangsan"
	fmt.Println("=========", stu.Name)
}

func main() {
	var stu = Student{"张三", 23}
	stu.show()
	fmt.Println("--------", stu)

	// 给结构体绑定的方法, 结构体指针也能调用, 只不过在调用show方法是传递的是结构体类型, 不是结构体指针类型
	var stu2 = &Student{"张三", 23}
	stu2.show()
	fmt.Print(stu2)
}
