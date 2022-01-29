package main

import (
	"fmt"
)

/**
* 1. if语句中esle不能换行
* 2. 只有一行语句时, if的{}不能省略
* 3. if中的()需要省略
* 4. if语句中可以声明变量
 */
func main() {
	// go语言中不建议这里写括号, age可以放在if语句中进行声明, 这是go的特殊用法
	if age := 23; age > 18 {
		fmt.Println("已经18岁")
	} else {
		fmt.Println("不满18岁")
	}
}
