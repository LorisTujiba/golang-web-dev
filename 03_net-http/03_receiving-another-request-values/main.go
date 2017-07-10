package main

/*
Exploring *http.request, we're going to try to
display the http request attributes
*/

import (
	"html/template"
	"net/http"
	"net/url"
)

var tpl *template.Template

type server int

func (s server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	//create a struct as a container
	data := struct {
		Method        string
		URL           *url.URL    // type URL from the net/url package
		Submissions   url.Values  //map of string
		Header        http.Header //its a string map
		Host          string
		ContentLength int64
	}{
		req.Method, //getting the method value form the http request
		req.URL,
		req.Form, //getting the form value from http request
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.Execute(w, data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var s server
	http.ListenAndServe(":8080", s)
}

//to run, run this
//then open browser
//submission will be nil if there's no data being passed
