package main

import "fmt"

func main() {
	var str = "你好, 世界"
	fmt.Println(str)

	// 反引号声明字符串(不用转移特殊字符)
	var str2 = `张三李四王五*\`
	fmt.Println(str2)

	// 更改字符串变量的引用
	str = "Hello Wrold"
	fmt.Println(str)

	// 不允许单独修改字符串
	// str[0] = 'H'
	// fmt.Println(str)

	// 拼接字符串, 注意换行的时候, 一定要把+保留在换行的后面, 这是因为golong会默认给每一行添加;结尾, 放发现行尾是+号时, 变不在添加
	var str3 = "Hello World" + "你好, 世界"
	fmt.Println(str3)
}