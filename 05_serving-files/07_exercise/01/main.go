package main

/*
ListenAndServe on port 8080 of localhost

For the default route "/" Have a func called "foo" which writes to the response "foo ran"

For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

This is from dog

" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
*/

import (
	"net/http"
	"io"
	"html/template"
	"os"
)

var tpl *template.Template

func foo (w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"foo ran")
}

func dog (w http.ResponseWriter,r *http.Request){
	tpl.Execute(w,"This is from dog")
}

func file (w http.ResponseWriter, r *http.Request){
	fl,err := os.Open("dog.jpg")
	if err != nil{
		http.Error(w,"Not found",404)
	}
	http.ServeFile(w,r,fl.Name())
}

func init(){
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/dog",dog)
	http.HandleFunc("/dog.jpg",file)
	http.ListenAndServe(":8080",nil)
}