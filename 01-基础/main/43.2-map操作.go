package main

import (
	"fmt"
)

func main() {
	var mymap = map[string]string{
		"city1": "北京",
		"city2": "上海",
		"city3": "深圳",
	}

	//添加数据
	mymap["city4"] = "濮阳"
	fmt.Println(mymap)

	// 查询数据
	fmt.Println(mymap["city1"])
	fmt.Println(mymap["ssss"])

	// 修改数据
	mymap["city2"] = "ShangHai"
	fmt.Println(mymap)

	// 删除数据(如果key不存才是不会报错的)
	delete(mymap, "city3")
	delete(mymap, "ssss")
	fmt.Println(mymap)

	// 查看map长度
	fmt.Println(len(mymap))

	// 删除所有的key: 一一遍历进行删除或者重新创建一个map值
}
