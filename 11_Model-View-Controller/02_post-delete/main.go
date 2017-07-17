package main

import (
	"encoding/json"
	"fmt"
	"github.com/LorisTujiba/golang-web-dev/11_Model-View-Controller/02_post-delete/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*=========================================
Model
=========================================
Create a model with json files. Save the
model in another package. try to
implement the MVC!
*/

/*=============================================================================================================================================
Curl
==============================================================================================================================================
From now on, lets execute things using curl. Start your server, then open another Git Bash and run these command below:

use curl -X POST -H "Content-Type: application/json" -d '{"Name":"James Bond","Gender":"male","Age":32,"Id":"777"}' http://localhost:8080/user
-X is short for --request Specifies a custom request method to use when communicating with the HTTP server.
-H is short for --header
-d is short for --data
curl -X DELETE -H "Content-Type: application/json" http://localhost:8080/user/777
*/

func main() {

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	r.POST("/user", createUser)
	r.DELETE("/delete", deleteUser)
	http.ListenAndServe(":8080", r)

}

func deleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//write delete code
	w.WriteHeader(http.StatusOK)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	//to demonstrate it, lets just change the id
	u.ID = "007"

	//re marshal
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Fprintln(w, "Welcome to Index!")

}

func getUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	u := models.User{
		Name:   "Loris",
		Gender: "Male",
		Age:    23,
		Id:     params.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
