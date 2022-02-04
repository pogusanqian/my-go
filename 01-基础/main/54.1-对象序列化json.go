package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string `json:"name"` //反射机制更改名字
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}

	//将monster序列化, 返回的data是一个byte[]
	data, _ := json.Marshal(&monster)
	fmt.Println(string(data))

}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, _ := json.Marshal(a)
	fmt.Println(string(data))
}

func testSlice() {
	// 切片中的元素是map
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	data, _ := json.Marshal(slice)
	fmt.Println(string(data))

}

//对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67
	data, _ := json.Marshal(num1)
	fmt.Println(string(data))
}

func main() {
	// 结构体序列化成json字符串
	testStruct()
	// map序列化成json字符串
	testMap()
	// 切片序列化成json
	testSlice()
	// 基本数据类型序列化
	testFloat64()
}
