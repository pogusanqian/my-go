package main

import (
	"fmt"
)

func main() {
	// 数组中套了一个map
	var mapSlice []map[string]string
	mapSlice = make([]map[string]string, 2)

	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]string, 2)
		mapSlice[0]["name"] = "北京"
		mapSlice[0]["pinyin"] = "BeiJing"
	}

	if mapSlice[1] == nil {
		mapSlice[1] = make(map[string]string, 2)
		mapSlice[1]["name"] = "上海"
		mapSlice[1]["pinyin"] = "ShangHai"
	}

	// 因为我们设计的mapSlice只有两个长度, 所以这里使用append()进行动态添加
	var city3 = map[string]string{
		"name":   "深圳",
		"pinyin": "ShenZhen",
	}
	mapSlice = append(mapSlice, city3)

	fmt.Println(mapSlice)
}
