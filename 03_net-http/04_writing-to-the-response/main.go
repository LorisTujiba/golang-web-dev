package main

import "net/http"

type server int

func (s server)ServeHTTP(w http.ResponseWriter,res *http.Request){
	w.Header().Set("Loris-Key","This is from me")//set the header of the response
	w.Header().Set("Content-Type","text/html;charset=utf-8")
}

func main(){
	var s server
	http.ListenAndServe(":8080",s)
}

//use web browser, hit f12,check the header