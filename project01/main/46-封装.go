package main

import (
	"fmt"
	"pogusanqian.com/project01/util"
)

/**
* 1. 封装就是将抽象出来的字段和对字段的操作封装在一起, 数据被保护在内部, 程序的其他包只有通过被授权的方法才能调用;
* 2. 封装可以实现隐藏细节, 以及校验数据保证数据安全性的特点
 */
func main() {
	var stu = util.NewStudent("张三", 23) // 现在返回的stu是一个指针
	stu.SetSex("男")
	fmt.Println(stu.GetSex())
}
