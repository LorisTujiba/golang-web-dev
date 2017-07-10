package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", handleHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	var s string
	if r.Method == http.MethodPost {

		//open
		f, h, err := r.FormFile("theFile") //catches the file, can be image, text etc
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		//fyi
		fmt.Println(f, " : ", h, " : ", err)

		//read
		bs, err := ioutil.ReadAll(f) //read the file, then get the byte slice
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs) //parse to string
	}

	err := tpl.Execute(w, s)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
