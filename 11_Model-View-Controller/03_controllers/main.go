package main

import (
	"github.com/LorisTujiba/golang-web-dev/11_Model-View-Controller/03_controllers/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*==========================================================================
Controller
==========================================================================
Now, lets implement the controllers. You see, now main.go looks very clean
because we separate things. remember, separation of concerns. Lets look
into our controllers/user.go
*/

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/", uc.Index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/delete", uc.DeleteUser)
	http.ListenAndServe(":8080", r)

}
