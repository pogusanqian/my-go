package main

import (
	"html/template"
	"net/http"
)

func tempTemplate(w http.ResponseWriter, r *http.Request) {
	// 设置两个html文件成文模板
	t := template.Must(template.ParseFiles("../static/tempTemplate1.html", "../static/tempTemplate2.html"))
	// 这里的张三只能传递到tem1中, 我们使用templete关键字传递到tem2中
	t.Execute(w, "张三")
}

func main() {
	http.HandleFunc("/tempTemplate", tempTemplate)
	http.ListenAndServe(":8080", nil)
}
