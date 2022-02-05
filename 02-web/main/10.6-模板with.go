package main

import (
	"html/template"
	"net/http"
)

func tempWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../static/tempWith.html"))
	t.Execute(w, "张三")
}

func main() {
	http.HandleFunc("/tempWith", tempWith)
	http.ListenAndServe(":8080", nil)
}
