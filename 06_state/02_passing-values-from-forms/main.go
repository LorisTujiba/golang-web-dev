package main

/*
Sometimes web browser gonna looking for an favicon.ico
so to handle that, we can use this code below
*/

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

type user struct {
	FirstName string
	LastName  string
}

func init() {
	tpl = template.Must(template.ParseFiles("example2.gohtml"))
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/2", handleHome2)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleHome2(w http.ResponseWriter, r *http.Request) {
	fName := r.FormValue("fname")
	lName := r.FormValue("lname")

	err := tpl.Execute(w, user{fName, lName})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("q")                   //get value of property q from url
	w.Header().Set("Content-Type", "text/html") //post, sent from the body, get, goes to url
	io.WriteString(w, `
	<form method = "post">
	<input type="text" name="q">
	<input type="submit">
	</form><br>`+value)
}
