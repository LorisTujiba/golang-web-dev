package main

import (
	"net/http"
	"github.com/LorisTujiba/golang-web-dev/14_more-mongodb/13_implementation-and-code-org/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Employees)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/json", controllers.ShowJSON)// << this is the material that we added
	http.ListenAndServe(":8080", nil)
}
