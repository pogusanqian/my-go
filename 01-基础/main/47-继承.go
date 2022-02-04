package main

import (
	"pogusanqian.com/project01/util"
)

/**
 * 1. go语言可以实现多继承
 * 2. 结构体在使用嵌套继承的时候, 会继承父类的所有属性, 公开和私有属性都会被继承; 当结构体和父类有相同的属性时, 就会查找按照就近原则查找对应的属性
 */
func main() {
	var midstu = util.MiddleStudent{}
	// 这里原本的写法应该是midstu.student.Name, 进行了简写(注意这里这个student应该要是大写)
	// midstu结构体其实是没有Name属性的, 然后编译器没有找到, 就会到student这个结构体中进行查找
	midstu.Name = "张三"
	midstu.Age = 23
	midstu.SetSex("女") // 继承了私有属性sex

	midstu.Show()
}
