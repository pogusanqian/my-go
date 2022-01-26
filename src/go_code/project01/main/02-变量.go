package main
import "fmt"

// 声明全局变量
var county = "中国"

func main() {
	// 1. 指定变量类型, 声明后如果不赋值, 则使用默认值, int的默认值是0
	var age1 int
	fmt.Println("age1:", age1)

	// 2.根据值自动判定类型
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
		name1 = "李四"
		school1 = "河南理工"
	)
	fmt.Println(name1, school1)

	fmt.Println(county)
}