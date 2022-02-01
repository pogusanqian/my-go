package util

import (
	"fmt"
)

type Student3 struct {
	Name string
	Age  int
}

// 注意这里嵌套继承的时候, 是可以使用指针进行继承的
type MiddleStudent3 struct {
	*Student3 // 使用指针进行继承
	Age       int
}

func (stu *Student3) Show() {
	fmt.Println(stu.Name, stu.Age)
}

// 如果有两个重名的方法, 在在会报错ambiguous selector stu.Name
func (stu *MiddleStudent3) Show() {
	fmt.Println(stu.Age)
}
