package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World", r.URL.Path)
}

func main() {
	// HandlerFunc是一个适配器, 通过类型转换让我们可以将普通的函数作为HTTP处理器使用; 内部自动实现了实现了ServeHTTP接口, 但是要求方法的参数必须是http.ResponseWriter和*http.Request
	// handle是一个处理器, 用来处理http请求使用; Mux多路复用器就是路由, 把具体的请求分配到具体的处理器上
	http.HandleFunc("/", handler)

	// ListenAndServe监听TCP地址addr, 并且会使用handler参数调用Serve函数处理接收到的连接, 端口地址默认是80; handler参数一般会设为nil，此时会使用DefaultServeMux。
	http.ListenAndServe(":8080", nil)
}
