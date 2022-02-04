package main

import (
	"pogusanqian.com/project01/util"
)

func main() {
	var midstu = util.MiddleStudent2{}
	// 因为是多继承, 中间的父类结构体是不能省略的; 如果这里不写中间结构体是会报错的ambiguous selector midstu.Name
	midstu.Student2.Name = "张三"
	midstu.People2.Age = 23
	// show方法其实是绑定在student2结构体上的, 但是我们是给People2进行的赋值, 所以show方法输出的属性都是默认值
	midstu.Show()          // 这个调用的MiddleStudent2的show方法
	midstu.Student2.Show() // 这个调用的是Student2的Show方法
}
