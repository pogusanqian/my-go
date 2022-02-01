package main

import (
	"fmt"
	"pogusanqian.com/project01/util"
)

func main() {
	// 一次性声明属性值
	var midstu = util.MiddleStudent2{
		util.People2{"zhangsan", 22},
		util.Student2{"张三", 23},
		222,
	}
	fmt.Println(midstu)
	midstu.Show()
	midstu.Student2.Show()

	var midstu2 = util.MiddleStudent2{
		util.People2{
			Name: "zhangsan",
			Age:  22}, // 这与这种写法, 22后面是没有跟,号的
		util.Student2{
			Name: "张三",
			Age:  23,
		},
		222,
	}
	fmt.Println(midstu2)

}
