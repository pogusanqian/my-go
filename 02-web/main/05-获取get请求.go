package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	// 这里的query并没有转换成对象, 而是一个原始的query字符串
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
}

func main() {
	http.HandleFunc("/getStudent", handler)
	http.ListenAndServe(":8080", nil)
}
