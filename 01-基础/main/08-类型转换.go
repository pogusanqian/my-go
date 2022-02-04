package main

import (
	"fmt"
	"strconv"
)

/**
 * 1. go不支持隐式类型转换, 需要显示转换
 * 2. 高精度转低精度, 可能会有精度丢失
 * 3. 如果转换数据失败是, go不会进行报错, 而是设置成一个默认值(就算原本由值, 也设置成默认值)
 */
func main() {
	// 数据类型长度转换
	var num int32 = 100
	var num2 int8 = int8(num)

	fmt.Printf("%T %v \n", num2, num2)

	// 基本类型转字符串
	var str string = fmt.Sprintf("%d", 123)
	var str2 string = fmt.Sprintf("%t", true)

	fmt.Printf("%T %q \n", str, str)
	fmt.Printf("%T %q \n", str2, str2)

	// strconvn将基本数据类型转换成字符串
	var str3 string = strconv.FormatInt(123, 10) // 10表示转换成10进制
	fmt.Printf("%T %q \n", str3, str3)

	// strconvn将字符串转换成基本数据类型ParseBool()会返回两个值, 第二个值时err对象, 这里进行了忽略
	var flag, _ = strconv.ParseBool("true")
	fmt.Printf("%T %v \n", flag, flag)

	// 转换失败, 设置默认值
	var falg = true
	falg, _ = strconv.ParseBool("hahah")
	fmt.Printf("%T %v \n", falg, falg)

}
