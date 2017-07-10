package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCook)
	http.HandleFunc("/read", readCook)
	http.HandleFunc("/expire", exp)
	http.ListenAndServe(":8080", nil)
}

func setCook(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: "Joseph",
	})
	fmt.Println(w, "Cookie is written, chek the browser, dev tools / application / cookies")
}

func readCook(w http.ResponseWriter, r *http.Request) {
	co, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/setCook", http.StatusSeeOther)
		return // dont forget to return from this function
	}
	fmt.Fprintln(w, "here is the cookie , ", co)
}

func exp(w http.ResponseWriter, r *http.Request) {
	co, err := r.Cookie("username")
	if err != nil {
		http.Redirect(w, r, "/setCook", http.StatusSeeOther)
		return
	}
	co.MaxAge = -1 //delete the cookie
	http.SetCookie(w, co)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
