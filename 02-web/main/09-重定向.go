package main

import (
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	//设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/redirect", redirect)
	http.ListenAndServe(":8080", nil)
}
