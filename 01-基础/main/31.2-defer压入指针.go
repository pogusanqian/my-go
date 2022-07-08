package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// TODO 指针指向了同一块内存地址, 但是值确实不一样的
func sum(num *int) {
	defer fmt.Println("===========defer1", *num, num) // 10 0xc0000aa058
	*num = *num + 100
	fmt.Println("===========sum函数执行", *num, num) // 110 0xc0000aa058
}

func changeAge(stu *Student) {
	defer fmt.Printf("%p %v %v %v\n", stu, stu, *stu, stu.Age) // 0xc000004078 &{张三 100} {张三 100} 100
	stu.Age = 100
	fmt.Printf("%p %v %v %v\n", stu, stu, *stu, stu.Age) // 0xc000004078 &{张三 100} {张三 23} 23
}

func main() {
	num := 10
	sum(&num)
	fmt.Println(num)

	stu := Student{"张三", 23}
	changeAge(&stu)
	fmt.Println(stu)
}
