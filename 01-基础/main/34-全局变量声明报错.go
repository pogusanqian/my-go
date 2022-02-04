package main

import (
	"fmt"
)

var age = 23 // 正常, 相当于在声明age属性时, 初始化值时23
// name := "Tom" // 报错, 此语句可以等价与var name string; name = "Tom"; 赋值语句只能写在代码块中

func main() {
	fmt.Println("==========", age)
}
