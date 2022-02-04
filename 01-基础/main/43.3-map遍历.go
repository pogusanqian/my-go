package main

import (
	"fmt"
)

func main() {
	var mymap = map[string]string{
		"city1": "北京",
		"city2": "上海",
		"city3": "深圳",
	}

	// map不能使用for循环遍历, 只能使用range遍历
	for key, value := range mymap {
		fmt.Println(key, value)
	}
}
