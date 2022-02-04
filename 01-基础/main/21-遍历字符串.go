package main

import "fmt"

func main() {
	var str1 = "Hello World"
	var str2 = "你好, 世界"

	// 普通遍历字符串
	for i := 0; i < len(str1); i++ {
		// str[i]获取的一个int数字, 需要将其转换成字符; 如果是中文的话这里是会出现乱码的
		fmt.Printf("%c", str1[i])
	}

	fmt.Println()
	// for-range遍历
	for index, val := range str1 {
		// 获取的value值仍然是一个数字, 需要转换成对应的字符
		fmt.Printf("index=%d, value=%c \n", index, val)
	}

	// 转换成切片遍历中文
	str3 := []rune(str2)
	for index, val := range str3 {
		fmt.Printf("%d %v %c \n", index, val, val)
	}
}
