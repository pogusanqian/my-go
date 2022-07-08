package main

import (
	"html/template"
	"net/http"
)

func tempMapRange(w http.ResponseWriter, r *http.Request) {
	var emps = map[string]string{
		"name1": "张三",
		"name2": "李四",
		"name3": "王五",
	}

	t := template.Must(template.ParseFiles("../static/tempMapRange.html"))
	t.Execute(w, emps)
}

func main() {
	http.HandleFunc("/tempMapRange", tempMapRange)
	http.ListenAndServe(":8080", nil)
}
