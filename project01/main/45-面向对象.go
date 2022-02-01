package main

import (
	"fmt"
)

/**
* 面向对象(OPP)
* 1. Go并不是纯粹的面向对象, 而是支持面向对象
* 2. Go中没有class类, 而是采用了struct结构体
* 3. Go的面向对象去掉了继承, 重载, 构造函数, 析构函数, 隐藏this指针等等; 但是Go仍然可以实现继承, 封装, 多态三大特性
* 4. 结构体是值类型, 默认是值copy
* 5. 结构体的所有字段在内存中是连续的
 */

// 定义Cat结构体, 结构体就是一种可以自定义的数据类型
type Cat struct {
	Name  string
	Age   byte
	Color string
}

func main() {
	// 第一种声明方式
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 1
	cat1.Color = "白色"
	fmt.Println(cat1, cat1.Name)

	// 第二种声明方式(根据顺序)
	var cat2 = Cat{"小黑", 3, "黑色"}
	fmt.Println(cat2)

	// 第三种声明方式, 蓝色后面的,号不能省略
	var cat3 = Cat{
		Age:   2,
		Name:  "小蓝",
		Color: "蓝色",
	}
	fmt.Println(cat3)

	// 第四种声明方式, 这里的cat4是一个指针, cat4是一个指针, 可以使用cat4.Coloe, Go引擎会将其转换成指针(*cat4).Color
	var cat4 *Cat = new(Cat)
	(*cat4).Name = "小灰" // 这里的()是不能省略的, 因为.的优先级要比*高
	(*cat4).Age = 3
	cat4.Color = "灰色"
	fmt.Println(*cat4)

	// 第五种方式
	var cat5 *Cat = &Cat{"小黄", 5, "黄色"}
	fmt.Println(*cat5)

}
