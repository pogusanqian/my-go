package main

import (
	"fmt"
)

func main() {
	myslice := []string{"张三", "李四", "王五"}

	// for循环遍历
	for i := 0; i < len(myslice); i++ {
		fmt.Println(myslice[i])
	}

	// range遍历
	for _, value := range myslice {
		fmt.Println(value)
	}
}
