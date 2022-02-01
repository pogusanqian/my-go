package main

import (
	"fmt"
)

/**
* 1. 类型断言: b := a(Student); 判断a是不是Student类型的变量, 如果是就将a转换成student类型, 并赋值给b; 如果a不是Student类型就会报错
* 2. 空接口interface{}可以接收任意的数据类型
 */
func main() {
	var x interface{}
	var num1 float32 = 1.5

	fmt.Println(x)
	x = num1 // 空接口可以接受任意的数据类型
	fmt.Println(x)

	// 类型断言
	if y, ok := x.(float64); ok == true { // x是不能转换成float类型的
		fmt.Println("%t, %v", y, y)
	} else {
		fmt.Println("类型转换失败", y)
	}
}
