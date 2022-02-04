package model

import (
	"fmt"
	"testing"
)

// // TestMain函数可以在测试函数执行之前做一些其他操作; 必须执行m.Run(), 否则后面的测试方法是不会执行的
// func TestMain(m *testing.M) {
// 	fmt.Println("测试开始：")
// 	m.Run()
// }

// // 使用t.Run()来执行子测试函数
// func TestUser(t *testing.T) {
// 	t.Run("测试添加用户:", add)
// }

// 如果函数名不是以Test开头, 那么该函数默认不执行, 但是如果这个函数是其他测试函数的子函数, 便会被调用
func add(t *testing.T) {
	stu := &Student{
		Name: "lisi",
		Age:  40,
		Sex:  "女",
	}
	stu.AddStudent()
	stu.AddStudent2()
}

func TestQuery(t *testing.T) {
	student := &Student{Id: 1}
	stu, _ := student.GetStudentByID()
	stus, _ := student.GetStudents()
	fmt.Println("---------", stu)
	fmt.Println("---------", stus)
}
