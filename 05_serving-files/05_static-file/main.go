package main

/*
Static file server for static website

This code below will show the dir

IF, we have file with the name of index of html, it will no longer shows the files on
the dir

commonly in static website, there has to be an index.html file

*/

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
