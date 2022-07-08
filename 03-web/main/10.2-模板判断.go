package main

import (
	"html/template"
	"net/http"
)

func temif(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../static/temif.html"))
	t.Execute(w, true)
}

func main() {
	http.HandleFunc("/temif", temif)
	http.ListenAndServe(":8080", nil)
}
