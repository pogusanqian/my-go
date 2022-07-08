package main

import (
	"html/template"
	"net/http"
)

type Employee struct {
	ID    int
	Name  string
	Email string
}

func tempRange(w http.ResponseWriter, r *http.Request) {
	var emps []*Employee
	emp := &Employee{
		ID:    1,
		Name:  "李小璐",
		Email: "lxl@jnl.com",
	}
	emps = append(emps, emp)
	emp2 := &Employee{
		ID:    2,
		Name:  "白百何",
		Email: "bbh@cyf.com",
	}
	emps = append(emps, emp2)
	emp3 := &Employee{
		ID:    3,
		Name:  "马蓉",
		Email: "mr@wbq.com",
	}
	emps = append(emps, emp3)

	t := template.Must(template.ParseFiles("../static/tempRange.html"))
	t.Execute(w, emps)
}

func main() {
	http.HandleFunc("/tempRange", tempRange)
	http.ListenAndServe(":8080", nil)
}
