package main

/*
Sometimes web browser gonna looking for an favicon.ico
so to handle that, we can use this code below
 */

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",handleHome)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func handleHome(w http.ResponseWriter,r *http.Request){
	value := r.FormValue("q")//get value of property q from url
	io.WriteString(w,"The value from the url is : "+value)
}

//run http://localhosy:8080/?q=dog
