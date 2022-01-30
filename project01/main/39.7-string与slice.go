package main

import (
	"fmt"
)

/**
* string底层是一个[]byte, 所以string也可以进行切片处理
 */
func main() {
	var str = "helloworld"
	var str2 = "你好世界"

	// 使用切片截取数组, 修改切片值会同步修改字符串值, 因为他们操作的是同一块空间, 但是由于字符串的值是不能改变的, 所以使用string的切片是不能修改值的
	myslice1 := str[:5]
	fmt.Println(str, myslice1)

	// 将字符串转换成byte数组进行修改(这里只能处理英文, 三个字节的中文是不能处理的)
	arr := []byte(str)
	arr[0] = 'H' // 这里只能使用''号的字符, 不能使用字符串
	fmt.Println(string(arr))

	// 转换成rune切片来处理中文数据
	arr2 := []rune(str2)
	arr2[0] = '中'
	fmt.Println(string(arr2))

}
