package main

/*
Getting values from form that sent from html
*/

import (
	"html/template"
	"log"
	"net/http"
)

type server int

var tpl *template.Template

func (s server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) //hget the values from the form
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var s server
	http.ListenAndServe(":8080", s) //if a request coming, handle with s
	//s will use the serveHTTP, and remember that s is a handler
}

//run this, open localhost:8080 from your browser
