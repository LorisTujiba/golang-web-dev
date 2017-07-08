package main

import (
	"net/http"
	"html/template"
)

/*
If you want to send file, put the multipart/form-data
the default enctype is 'application/x-www-form-urlenconded'
urlencoded will return a key value pairs separated by &

the last type of enctype is text/plain
 */


var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/",dHandle)
	http.HandleFunc("/textplain",textplainHandle)
	http.HandleFunc("/multihandle",multiHandle)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func textplainHandle(w http.ResponseWriter,r *http.Request){

	bs := make([]byte,r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w,"textplain.gohtml",body)//try each of the html
	if err != nil{
		http.Error(w,err.Error(),500)
	}
}

func multiHandle(w http.ResponseWriter,r *http.Request){

	bs := make([]byte,r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w,"index.gohtml",body)//try each of the html
	if err != nil{
		http.Error(w,err.Error(),500)
	}
}

func dHandle(w http.ResponseWriter,r *http.Request){

	bs := make([]byte,r.ContentLength)//biar praktis ambil dari sini aja, gausa cape bikin form lol
	r.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w,"default.gohtml",body)//try each of the html
	if err != nil{
		http.Error(w,err.Error(),500)
	}
}