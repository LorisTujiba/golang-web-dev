package main

import (
	"net/http"
	"fmt"
)

/*
StatusMovedPermanently = 301 // RFC 7231, 6.4.2
StatusSeeOther         = 303 // RFC 7231, 6.4.4 specific, always change the method to get
StatusTemporaryRedirect = 307 // RFC 7231, 6.4.7 preserve the method
*/

func main(){
	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter,r *http.Request){
	fmt.Println("home - req method : ",r.Method,"\n\n")
}
func bar(w http.ResponseWriter,r *http.Request){
	fmt.Println("bar - req method : ",r.Method,"\n\n")
	http.Redirect(w,r,"/",http.StatusMovedPermanently)
}

/*
if you go to bar, first time get
the second time, you can no longer to bar
 */

