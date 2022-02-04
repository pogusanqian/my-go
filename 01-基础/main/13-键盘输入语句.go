package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32

	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水:")
	fmt.Scanln(&sal)

	fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v", name, age, sal)
}
