package main

import (
	"html/template"
	"net/http"
)

/*Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc
Contstraint: Do not change anything outside of func main

Hints:

http.HandlerFunc

type HandlerFunc func(ResponseWriter, *Request)
http.HandleFunc

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
source code for HandleFunc

  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  		mux.Handle(pattern, HandlerFunc(handler))
  }
*/

var tpl *template.Template
var name = "Loris Tujiba Soejonopoetro"

func handleIndex(w http.ResponseWriter, response *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func handleDog(w http.ResponseWriter, response *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "Dog page")
}

func handleMe(w http.ResponseWriter, response *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", name)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/", http.HandlerFunc(handleIndex))
	http.Handle("/dog", http.HandlerFunc(handleDog))
	http.Handle("/me", http.HandlerFunc(handleMe))
	http.ListenAndServe(":8080", nil)

}
