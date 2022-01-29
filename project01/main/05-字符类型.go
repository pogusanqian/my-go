package main

import "fmt"

/**
* 1. 字符类型使用单引号''
* 2. go语言使用的编码是Utf-8
* 3. go语言中没有专门的数据类型来存储字符类型, 如果存储字符类型的话, 可以使用byte类型保存(如果溢出, 可以使用int存储)
 */
func main() {
	// 声明字符类型
	var c byte = 'a'
	fmt.Println(c)

	// 溢出处理
	var b int = '北'
	fmt.Println(b)

	// 将97转换成对应的字符输出
	fmt.Printf("%c", 97)
	fmt.Println()

	// 字符运算, 这里写成"a" + 10就会报错
	var res1 = 'a' + 10
	fmt.Println(res1)
}
