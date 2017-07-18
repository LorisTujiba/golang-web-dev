package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/03_exercise/02/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
)

//UserController struct, saves a user data model
type UserController struct {
	data map[string]models.User
}

//NewUserController initiate, returns an address
func NewUserController(u map[string]models.User) *UserController {
	return &UserController{u}
}

//GetUser , getting a user from the model
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	user := uc.data[id]

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

//CreateUser , store user data to the model
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	u.ID = uuid.NewV4().String()

	//store user
	uc.data[u.ID] = u

	uj, _ := json.Marshal(u)
	models.StoreUsers(uc.data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

//DeleteUser , delete the data from the storage
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.data, id)
	models.StoreUsers(uc.data)

	w.WriteHeader(http.StatusOK) // 200
}
