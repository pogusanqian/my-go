package main

import (
	"html/template"
	"net/http"
)

func tempSelectDifne(w http.ResponseWriter, r *http.Request) {
	// 这里传递的是tempContent1模板, 就会加载这个文件的模板; 如果是tempContent1.html, 就会加载tempContent1中的模板; 不传递的话会使用默认模板
	t := template.Must(template.ParseFiles("../static/tempSelectDifne.html", "../static/tempContent1.html"))
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/tempSelectDifne", tempSelectDifne)
	http.ListenAndServe(":8080", nil)
}
