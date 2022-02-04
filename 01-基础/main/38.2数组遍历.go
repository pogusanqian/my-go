package main

import (
	"fmt"
)

func main() {
	var arr = [3]string{2: "王五", 0: "张三", 1: "李四"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//index如果不想使用, 可以使用_忽略
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
