package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	hello "hello/pb"
)

func main() {
	sex := "男"
	stu := &hello.Student{
		Name:   "张三",
		Age:    23,
		Sex:    &sex,
		School: []string{"油田一中", "河南理工大学"},
	}
	fmt.Println(stu)

	// 序列化
	marshal, err := proto.Marshal(stu)
	if err != nil {
		panic(err)
	}
	fmt.Println(marshal)

	// 反序列化
	newStu := &hello.Student{}
	err = proto.Unmarshal(marshal, newStu)
	if err != nil {
		panic(err)
	}
	fmt.Println(newStu)
}
