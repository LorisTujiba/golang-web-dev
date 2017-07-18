package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/02_CRUD-with-mongodb/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//UserController create an empty interface
type UserController struct {
	session *mgo.Session //create the session for the controller
}

//NewUserController , create a func that point to the interface
func NewUserController(s *mgo.Session) *UserController { //takes a session
	return &UserController{s}
}

//DeleteUser from mongodb
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//CreateUser store suer data to mongodb
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u) //decode the request body

	//create bson ID
	u.ID = bson.NewObjectId()

	//store the user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

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

//GetUser get user from the mongodb
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	//get the param first
	id := params.ByName("id")

	//verify id is objectid hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	//objectidhex returns as object id from the provided hex representation
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	//fetch user
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
