package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("."))) //display what file there's on the directory
	http.HandleFunc("/loris", handle)
	http.ListenAndServe(":8080", nil)

}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="loris.jpg">`)
}
