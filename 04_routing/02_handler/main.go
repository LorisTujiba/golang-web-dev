package main

import (
	"io"
	"net/http"
)

//type HandlerFunc is different with a function that asked responseWriter and *Request

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "C C C C C")
}

func main() {

	http.HandleFunc("/c", c) //totally works
	//http.Handle("/c2",c) how to make this works?
	http.Handle("/c2", http.HandlerFunc(c)) //convert it to handlerFunc
	//Handlerfunc has ServeHTTP attached to it,
	//which makes it a handler, and that's what Handle want

	http.ListenAndServe(":8080", nil) //use the default serve mux
}
