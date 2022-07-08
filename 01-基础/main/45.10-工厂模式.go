package main

import (
	"fmt"
	"pogusanqian.com/project01/util"
)

/**
* 1. Go语言中是没有构造函数, 可以使用工厂模式创建首字母是小写的结构体
 */
func main() {
	var stu = util.NewStudent("张三", 23) // 现在返回的stu是一个指针
	fmt.Println(*stu)
}
