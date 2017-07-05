package main

import (
	"net/http"
	"io"
)

type server int

func (s server)ServeHTTP(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"ser ser server")
}

type newfoundland int

func (s newfoundland)ServeHTTP(w http.ResponseWriter,req *http.Request){
	io.WriteString(w,"Newfoundland")
}

func another (w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "HandleFunc wants a function with response writer and " +
		"*http request as the signature. As long the function have the signature" +
		"it's okay")
}

func main(){

	var d server
	var n newfoundland

	mux := http.NewServeMux()//Create the mux
	mux.Handle("/",d)//if the request path is /dog, then d gonna handle
	mux.Handle("/newfoundland",n)//otherwise, if the request path is <, then n gonna handle the request
	mux.HandleFunc("/another",another)

	http.ListenAndServe(":8080",mux)// pass the mux as the handler type
}
