package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32

	fmt.Println("请输入姓名, 年龄, 薪水, 使用空格隔开")
	// 如果输入的不符合格式, 会给设置一个默认值
	fmt.Scanf("%s %d %f", &name, &age, &sal)

	fmt.Printf("姓名: %v, 年龄: %v, 薪水: %v", name, age, sal)
}
