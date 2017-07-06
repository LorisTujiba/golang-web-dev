package main

import (
	"net/http"
	"os"
	"io"
)

/*
Let's  serve an image using the io.copy command
 */

func notBeingServedFromOurServer(w http.ResponseWriter,r *http.Request){

	w.Header().Set("Content-Type","text/html;charset=utf-8")
	io.WriteString(w,`<img src="loris.jpg">`)
}

func handle(w http.ResponseWriter, r *http.Request){

	fl,err := os.Open("loris.jpg")
	if err!=nil{
		http.Error(w,"File Not Found",404)//returns http error 404
		return
	}
	defer fl.Close()

	io.Copy(w,fl)

}

func main(){

	http.HandleFunc("/",notBeingServedFromOurServer)
	http.HandleFunc("/loris.jpg",handle)
	http.ListenAndServe(":8080",nil)

}
