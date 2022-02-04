package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 格式化参数, 会将parmas解析到r.From属性中, 表单参数解析到r.PostForm属性中
	// ParseForm只支持格式化普通的form表单, 如果是上传的form表单需要使用MultipartForm()
	// 两外还有两个方法FormValue(), 和PostFormValue()会隐式调用ParseMultipartForm和ParseForm两个函数
	r.ParseForm()
	//获取请求参数
	fmt.Fprintln(w, "请求参数有：", r.Form)
	fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("id"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("name"))
}

func main() {
	http.HandleFunc("/parseForm", handler)
	http.ListenAndServe(":8080", nil)
}
