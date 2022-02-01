package util

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
	sex  string // 这个属性特意写的小写, 为了实现属性的封装
}

type MiddleStudent struct {
	student     // 继承student方法
	Age     int // 添加一个Age属性, 如果在使用MiddleStudent调用show方法的时候, 其实输出的Age是0, 因为值并没有添加到结构体student上, 反而是添加到了MiddleStudent上
}

// 使用工厂模式创建函数, 注意这里返回的是一个指针类型, 这样就不是值copy了
func NewStudent(name string, age int) *student {
	return &student{
		Name: name,
		Age:  age,
	}
}

// 给student结构体添加方法
func (stu *student) SetSex(sex string) {
	if sex != "男" && sex != "女" {
		fmt.Println("性别只能是男女")
	} else {
		stu.sex = sex
	}
}

func (stu *student) GetSex() string {
	return stu.sex
}

func (stu *student) Show() {
	fmt.Println(stu.Name, stu.Age, stu.sex)
}
