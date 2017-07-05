package main

import (
	"net/http"
	"io"
)

/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

var name = "Loris Tujiba Soejonopoetro"

func handleIndex(w http.ResponseWriter,response *http.Request){
	io.WriteString(w,"This is an index page")
}

func handleDog(w http.ResponseWriter,response *http.Request){
	io.WriteString(w,"This is a dog page")
}

func handleMe(w http.ResponseWriter,response *http.Request){
	io.WriteString(w,name)
}

func main(){

	http.HandleFunc("/",handleIndex)
	http.HandleFunc("/dog",handleDog)
	http.HandleFunc("/me",handleMe)
	http.ListenAndServe(":8080",nil)

}
