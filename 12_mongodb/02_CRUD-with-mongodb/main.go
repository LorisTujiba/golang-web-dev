package main

/*===============================================
BSON
===============================================
We're going to use BSON(Binary version of JSON)
check the models

to run, remember, try and learn to use curl
comfy approach. Try to see the ID, cuz we
are using the bson type :)

======================
POST a user to mongodb
======================
Enter this at the terminal

curl -X POST -H "Content-Type: application/json" -d '{"username":"James Bond","gender":"male","age":32}' http://localhost:8080/user

-X is short for --request Specifies a custom request method to use when communicating with the HTTP server.
-H is short for --header
-d is short for --data

=======================
GET a user from mongodb
=======================

Enter this at the terminal

curl http://localhost:8080/user/<enter-user-id-here>

==========================
DELETE a user from mongodb
==========================
Enter this at the terminal

curl -X POST -H "Content-Type: application/json" -d '{"username":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
curl http://localhost:8080/user/<enter-user-id-here>

curl -X DELETE http://localhost:8080/user/<enter-user-id-here>


*/

import (
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/02_CRUD-with-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/delete/:id", uc.DeleteUser)
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
