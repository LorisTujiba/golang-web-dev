package main

import (
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/01_setup-and-connect/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

/*=========================================================
MongoDB
=========================================================
MongoDB is a document storage, we can just store JSON
inside it.

First, go get the mongodb driver

go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson

So, we will need a mongo session to use in the CRUD methods
we need the Controllers to have access to a mongo session

Let's add this to controllers/user.go

UserController struct {
    session *mgo.Session
}

and add a param for the newusercontroller that accept the mgo session

And now add this to main.go
then create a getsession function

Don't forget to start mongo db

Enter this at the terminal
curl http://localhost:8080/user/1
*/

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/delete", uc.DeleteUser)
	http.ListenAndServe(":8080", r)

}

func getSession() *mgo.Session {
	//connect to local mongo
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
