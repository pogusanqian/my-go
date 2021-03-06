package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//1. 统计字符串的长度, 按字节 len(str), golang的编码统一为utf-8 (ascii的字符(字母和数字) 占一个字节, 汉字占用3个字节)
	str := "hello北"
	var str2 string
	fmt.Println("str len=", len(str)) // 8

	//1. 字符串转整数:	 n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	//3. 整数转字符串  str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	//4. 字符串 转 []byte:  var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//5. []byte 转 字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//6. 10进制转 2, 8, 16进制:  str = strconv.FormatInt(123, 2),返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)

	//7. 查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//8. 统计一个字符串有几个指定的子串
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v\n", num)

	//9. 不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)           //true
	fmt.Println("结果", "abc" == "Abc") // false //区分字母大小写

	//10.返回子串在字符串第一次出现的index值, 如果没有返回-1
	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v\n", index)

	//11. 返回子串在字符串最后一次出现的index, 如没有返回-1
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n", index)

	//12. 将指定的子串替换成另外一个子串, n可以指定你希望替换几个, 如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	//13. 按照指定的某个字符, 为分割标识, 将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,wrold,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	//14. 将字符串的字母进行大小写的转换
	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str) //golang hello

	//15. 将字符串左右两边的空格去掉
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str=%q\n", str)

	//16. 将字符串左右两边指定的字符去掉
	str = strings.Trim("! he!llo !", "!")
	fmt.Printf("str=%q\n", str)

	//17. 判断字符串是否以指定的字符串开头:
	b = strings.HasPrefix("ftp://192.168.10.1", "hsp")
	fmt.Printf("b=%v\n", b)
}
