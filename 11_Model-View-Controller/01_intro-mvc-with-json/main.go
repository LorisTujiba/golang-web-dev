package main

/*
Get Julien's Router first

and

try to run things using curl
open git bash, and run curl http://localhost:8080/user/2

====================================================================
MVC
==================================================================
MVC is a design pattern that organize your code based on 3 things.

1.Model, Where you store your data/storage/db model.
2.View, Where you store things that user will view like html,css
3.Controller, Where you store your web app logic flow

The purpose is to make it easier while debugging, because
you separate things. Separation of concerns, a structure
for keeping display and data separate to allow
each to change without affecting the other.
*/

import (
	"encoding/json"
	"fmt"
	"github.com/LorisTujiba/golang-web-dev/11_Model-View-Controller/01_intro-mvc-with-json/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New() //julien's router
	r.GET("/", index)
	// added route plus parameter, using julien's router
	r.GET("/user/:id", getUser) // :id the id is the param
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "go to /user/12123")
}

//to demonstrate it, pretend that we have already have the data
//and lets say you want to get user datas that has
//the param id
func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//create the model
	u := models.User{
		Name:   "Loris",
		Gender: "Male",
		Age:    23,
		ID:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write content-type, stat code,payload, and display it
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
