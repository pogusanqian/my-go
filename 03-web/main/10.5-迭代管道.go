package main

import (
	"html/template"
	"net/http"
)

func tempChannelRange(w http.ResponseWriter, r *http.Request) {
	strChan := make(chan string, 3)
	strChan <- "张三"
	strChan <- "李四"
	strChan <- "王五"
	close(strChan)

	t := template.Must(template.ParseFiles("../static/tempChannelRange.html"))
	t.Execute(w, strChan)
}

func main() {
	http.HandleFunc("/tempChannelRange", tempChannelRange)
	http.ListenAndServe(":8080", nil)
}
