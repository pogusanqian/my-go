package main

import (
	"html/template"
	"net/http"
)

func tempDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../static/tempDefine.html",))
	t.Execute(w, "张三")
}

func main() {
	http.HandleFunc("/tempDefine", tempDefine)
	http.ListenAndServe(":8080", nil)
}
