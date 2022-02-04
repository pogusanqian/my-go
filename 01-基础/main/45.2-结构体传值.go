package main

import (
	"fmt"
)

type Cat struct {
	Name  string
	Age   byte
	Color string
}

func main() {
	var cat = Cat{"小黑", 3, "黑色"}
	cat2 := cat // 结构体传递值时, 是值的copy
	cat2.Name = "大黑"

	fmt.Println(cat)  // {小黑 3 黑色}
	fmt.Println(cat2) // {小黑 3 黑色}
}
