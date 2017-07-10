package main

/*
Sometimes web browser gonna looking for an favicon.ico
so to handle that, we can use this code below
*/

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "go look at your terminal")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
