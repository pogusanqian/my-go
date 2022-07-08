package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name string `json: "name"`
	Age  int    `json: age`
}

func getString(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("我是字符串"))
}

func getJSON(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 将对象转换成json字符串进行传递
	json, _ := json.Marshal(Student{"张三", 23})
	w.Write(json)
}

func main() {
	http.HandleFunc("/getString", getString)
	http.HandleFunc("/getJSON", getJSON)
	http.ListenAndServe(":8080", nil)
}
