package main

import (
	"flag"
	"fmt"
)

var age int
var name = flag.String("name", "zhangsan", "姓名")

func init() {
	flag.IntVar(&age, "age", 1234, "年龄")
}

// go run .\main.go -name lisi -age 23 男 河南理工大学
func main() {
	// 将命令行中的参数值解析到变量上
	flag.Parse()

	fmt.Println("name: ", *name)
	fmt.Println("age: ", age)

	// 获取剩余的参数
	for index, val := range flag.Args() {
		fmt.Println(index, val)
	}
}
