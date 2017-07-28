package main

import (
	"net/http"
	"github.com/LorisTujiba/golang-web-dev/14_more-mongodb/14_mongo-and-go/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Employees)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/json", controllers.ShowJSON)
	http.ListenAndServe(":8080", nil)
}
