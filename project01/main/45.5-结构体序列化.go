package main

import (
	"encoding/json"
	"fmt"
)

// 打上tag之后, 序列化成字符串的时候, Name会编程name
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var stu = Student{"张三", 23}
	var str, _ = json.Marshal(stu) // marshal会使用反射
	fmt.Println(string(str))
}
