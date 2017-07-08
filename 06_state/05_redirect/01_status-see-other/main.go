package _1_status_see_other

import (
	"html/template"
	"net/http"
	"fmt"
)

/*
Client makes a request to a server, then that request is for certain resource
and then the server just redirect the client to another location

couple of different redirect
200s are the successes
300s are the redirects

https://golang.org/pkg/net/http/

StatusMovedPermanently = 301 // RFC 7231, 6.4.2
StatusSeeOther         = 303 // RFC 7231, 6.4.4 specific, always change the method to get
StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7 preserve the method

			-Request -
			-req line-
			-header  -
Client	->	-body    - -> server
	|						|
	<-	<-	-response-		<-
			-statline-
			-header  -
			-body    -

req line gonna have the method , uri, http version
statline gonna have http ver, stat code and reason phrase

example

req line
method		uri		http ver
get      	/dog  	http/1.1
post		/apply	http/1.1

stat line

http ver	statcode	reason phrase
http/1.1	200			ok

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
	w.WriteHeader(http.StatusSeeOther)//303, remember 303 will always sete the method to get
}

/*
if you go to barred,pass something

the return will be post and get, cuz 303
 */

func barred(w http.ResponseWriter,r *http.Request){
	fmt.Println("barred - req method : ",r.Method,"\n\n")
	tpl.Execute(w,nil)
}

