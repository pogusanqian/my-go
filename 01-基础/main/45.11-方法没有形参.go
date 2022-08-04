package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

// 没有使用形参接口
func (*Student) show(string) {
	fmt.Println("==============")
}

func main() {
	var stu = Student{"张三", 23}
	stu.show("Hello") // 这里的hello必须传递
}
