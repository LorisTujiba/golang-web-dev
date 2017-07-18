package main

/*
Use a map

Use the code in the previous folder "07_solution".

There is a map that holds all of the user data.

Every time a user is created or deleted, write this map as JSON to a file.

Also, when your program starts, if there is a file with JSON data in it, load that data.

IMPORTANT: Make sure you update your import statements to import packages from the correct location!
*/

/*
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
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/03_exercise/02/controllers"
	"github.com/LorisTujiba/golang-web-dev/12_mongodb/03_exercise/02/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/delete/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
