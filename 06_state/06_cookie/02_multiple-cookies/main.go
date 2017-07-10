package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCook)
	http.HandleFunc("/read", readCook)
	http.HandleFunc("/abundance", abundance)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "here is the cookie , ", co)
	co2, err := r.Cookie("general")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "here is the cookie , ", co2)
	co3, err := r.Cookie("specific")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "here is the cookie , ", co3)
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "gggg",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "ssss",
	})
	fmt.Println(w, "Cookie is written, chek the browser, dev tools / application / cookies")
}
