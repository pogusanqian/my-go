package main

import (
	"fmt"
)

type Student struct {
}

// 判断参数数据类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, x)
		}
	}
}

func main() {
	var n1 int32 = 30
	var n2 = 300
	var name = "tom"
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, name,  stu1, stu2)
}
