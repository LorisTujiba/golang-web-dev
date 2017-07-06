package main

import (
"net/http"
"io"
)

func main(){

	//strip prefix returns a handler, and that's what Handle() needed
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("./assets"))))//strip prefix gonna strip #2
	http.HandleFunc("/loris",handle)
	http.ListenAndServe(":8080",nil)

}

func handle(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`<img src="/resources/loris.jpg">`)//#2, strip the /resources/, so we all left with the loris.jpg
}