package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/data.txt", http.FileServer(http.Dir(".")))
	http.Handle("/test.gohtml", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
