package main

import (
	"html/template"
	"net/http"
)

func template1(w http.ResponseWriter, r *http.Request) {
	// 将html解析成模板对象, 这里有一个异常我们没有做处理, 可以使用Must()包裹起来进行处理; 如果写成相对路径的话, 需要再main目录
	t, _ := template.ParseFiles("D:/Code/my-go/02-web/static/tem1.html")
	// execute只能给一个模板赋值, 但是parseFiles()可以一次性将多个文件转换成一个模板对象
	t.Execute(w, "张三")
}

func template2(w http.ResponseWriter, r *http.Request) {
	// 通过Must函数让Go帮我们自动处理异常, 一次性加载两个html文件形成模板对象
	t := template.Must(template.ParseFiles("../static/tem1.html", "../static/tem2.html"))
	// 将值填写到tem2.html中, 返回的也是tem2.html
	t.ExecuteTemplate(w, "tem1.html", "张三")
}

func main() {
	http.HandleFunc("/template1", template1)
	http.HandleFunc("/template2", template2)
	http.ListenAndServe(":8080", nil)
}
