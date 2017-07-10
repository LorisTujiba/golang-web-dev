package main

import (
	"io"
	"net/http"
	"os"
)

/*
Let's  serve an image using the serve file command
*/

func notBeingServedFromOurServer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `<img src="loris.jpg">`) //not from our server, served from somewhere
}

func handle(w http.ResponseWriter, r *http.Request) {

	fl, err := os.Open("loris.jpg")
	if err != nil {
		http.Error(w, "File Not Found", 404) //returns http error 404
		return
	}
	defer fl.Close()

	fi, err := fl.Stat() //get the attributes data
	if err != nil {
		http.Error(w, "File Not Found", 404) //returns http error 404
		return
	}

	http.ServeContent(w, r, fi.Name(), fi.ModTime(), fl) //so we can access the name, and the last modified time
	http.ServeFile(w, r, "loris.jpg")                    //much simpler with this

}

func main() {

	http.HandleFunc("/", notBeingServedFromOurServer)
	http.HandleFunc("/loris.jpg", handle)
	http.ListenAndServe(":8080", nil)

}
