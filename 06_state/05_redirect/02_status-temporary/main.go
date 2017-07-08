package main

import (
	"html/template"
	"net/http"
	"fmt"
)

/*
StatusMovedPermanently = 301 // RFC 7231, 6.4.2
StatusSeeOther         = 303 // RFC 7231, 6.4.4 specific, always change the method to get
StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7 preserve the method
*/

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.HandleFunc("/barred",barred)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter,r *http.Request){
	fmt.Println("home - req method : ",r.Method,"\n\n")
}
func bar(w http.ResponseWriter,r *http.Request){
	fmt.Println("bar - req method : ",r.Method,"\n\n")
	//process form data
	w.Header().Set("Location","/") //set req line location to be /
	w.WriteHeader(http.StatusTemporaryRedirect)//307, preserve the method
}

/*
if you go to barred,pass something

the return will be post and post, cuz 301
 */

func barred(w http.ResponseWriter,r *http.Request){
	fmt.Println("barred - req method : ",r.Method,"\n\n")
	tpl.Execute(w,nil)
}

