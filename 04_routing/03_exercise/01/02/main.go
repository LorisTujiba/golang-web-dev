package main

import (
	"net/http"
	"html/template"
)

/*
Take the previous program in the previous folder and change it so that:
a template is parsed and served
you pass data into the template
*/

var tpl *template.Template
var name = "Loris Tujiba Soejonopoetro"

func handleIndex(w http.ResponseWriter,response *http.Request){
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}

func handleDog(w http.ResponseWriter,response *http.Request){
	tpl.ExecuteTemplate(w,"dog.gohtml","Dog page")
}

func handleMe(w http.ResponseWriter,response *http.Request){
	tpl.ExecuteTemplate(w,"me.gohtml",name)
}

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){

	http.HandleFunc("/",handleIndex)
	http.HandleFunc("/dog",handleDog)
	http.HandleFunc("/me",handleMe)
	http.ListenAndServe(":8080",nil)

}
