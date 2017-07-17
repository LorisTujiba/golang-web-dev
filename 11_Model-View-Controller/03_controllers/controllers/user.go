package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/LorisTujiba/golang-web-dev/11_Model-View-Controller/03_controllers/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//UserController create an empty interface
type UserController struct{}

//NewUserController , create a func that point to the interface
func NewUserController() *UserController {
	return &UserController{}
}

//now, lets make all these function implementing our interface

//DeleteUser to be coded
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//write delete code
	w.WriteHeader(http.StatusOK)
}

//CreateUser change the id
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	//change id
	u.ID = "007"

	//remarshal
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

//Index handle the index
func (uc UserController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Fprintln(w, "Welcome to Index!")

}

//GetUser pretend that we're getting a user data with the param's id
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	u := models.User{
		Name:   "Loris",
		Gender: "Male",
		Age:    23,
		ID:     params.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
