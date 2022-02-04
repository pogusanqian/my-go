package main

import "fmt"

/**
变量注意事项:
	1. 如果变量声明了类型, 可以改变值, 但是不能改变数据类型(JS是弱类型语言, 既能改变值, 又能改变数据类型)
	2. 声明变量后, 如果没有赋值, go会自动添加默认值(JS中是undefined)
	3. GO使用var声明的变量是块级作用域(JS中使用var声明的是函数级作用域)
*/
// 声明全局变量
var county = "中国"

func main() {
	// 1. 指定变量类型, 声明后如果不赋值, 则使用默认值, int的默认值是0
	var age1 int
	fmt.Println("age1:", age1)

	// 2.根据值自动判定类型, 也可以写成var age2 int = 25;
	var age2 = 25
	fmt.Println("age2:", age2)

	// 3.省略var, 使用:=; 相当于是先var age3 int, 然后age3=23
	age3 := 23
	fmt.Println("age3:", age3)

	// 4.多变量声明, 字符串的默认值是""
	var name, age, school string
	name = "张三"
	age = "23"
	school = "郑州大学"
	fmt.Println(name, age, school)

	// 5.多变量声明
	var (
		name1   = "李四"
		school1 = "河南理工"
	)
	fmt.Println(name1, school1)

	fmt.Println(county)
}
