package main

/*
It's actually ok to separate things , as long
we put them in the same folder

to run it, use

go run *.go

instead of

go run main.go

*/

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	http.SetCookie(w, cookie)

	// if the user exists already, get user
	var u user
	if username, ok := association[cookie.Value]; ok {
		u = userDatas[username]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session") //get the cookie
	if err != nil {                      //if there's already a cookie, redirect
		return false
	}
	username := association[cookie.Value]
	_, ok := userDatas[username] //check if the cookie contain the username, then use ok idiom
	return ok
}
