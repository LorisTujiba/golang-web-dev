package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	cook, err := r.Cookie("tracker")

	if err == http.ErrNoCookie {
		cook = &http.Cookie{
			Name:  "tracker",
			Value: strconv.Itoa(0),
		}
	}

	iteration := cook.Value

	number, err := strconv.Atoi(iteration)

	number++

	cook.Value = strconv.Itoa(number)
	http.SetCookie(w, cook)

	fmt.Println("How many times ", cook.Value)

}
