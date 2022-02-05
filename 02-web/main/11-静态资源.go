package main

import (
	"net/http"
)

func main() {
	// 配置静态服务器路由, "/static/"不能写成/static, 因为表示的是一个目录, 另外需要再命令行启动, 不要使用code runner启动
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../static"))))
	http.ListenAndServe(":8080", nil)
}
