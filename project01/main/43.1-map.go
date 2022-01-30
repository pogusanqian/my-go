package main

import (
	"fmt"
)

/**
* 1. map数据类型的key可以是int, string, bool, 指针, channel, 接口, 管道, 结构体, 但是不能是slice, map, function, 因为这三种数据类型不能使用==判断
* 2. map声明之后是不会分配内存的, 需要make之后才能分配内存; 数组声明之后是可以分配内存的
* 3. map的key是不按照我们输入的顺序进行排序的(而是会自动排序)
* 4. map会动态添加键值对
 */
func main() {
	// 第一个string表示的key的数据类型, 第二个string表示的value的数据类型
	var mymap map[string]string
	mymap = make(map[string]string, 2) // 分配内存
	mymap["no2"] = "李四"
	mymap["no1"] = "张三"
	mymap["no3"] = "王五"
	mymap["no1"] = "zhagnsan"
	fmt.Println(mymap)

	// 第二种声明方式
	var mymap2 = make(map[string]string)
	mymap2["city1"] = "北京"
	mymap2["city2"] = "深圳"
	fmt.Println(mymap2)

	// 第三种声明方式(上海后面的逗号不能省略)
	var mymap3 = map[string]string {
		"city1": "北京",
		"city2": "上海",
	}
	fmt.Println(mymap3)
}
