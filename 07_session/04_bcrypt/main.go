package main

/*
	For security purposes, we ofc need to hash the password
	,so to do that, we can use bcrypt package

	get this package first
	go get golang.org/x/crypto/bcrypt

	for now, we're only going to apply it to the sign up

	for the next lesson we're going to apply it on the
	login feature
*/

import (
	"github.com/LorisTujiba/gotraining/src/github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var tpl *template.Template

/*
Because we're going to use bcrypt, so we're going to change
the password type to slice of byte
*/

type user struct {
	Username string
	Password []byte
}

var userDatas = map[string]user{}
var association = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signUp)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	user := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, r *http.Request) {

	user := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "home.gohtml", user)

}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) { // if user already logged in , redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		//check if the username is available
		if _, ok := userDatas[username]; ok {
			http.Error(w, "Already Exists", http.StatusForbidden)
			return
		}

		SID := uuid.NewV4()

		cookie := &http.Cookie{
			Name:  "session",
			Value: SID.String(),
		}
		http.SetCookie(w, cookie)

		association[cookie.Value] = username

		//use bcrypt
		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userDatas[username] = user{username, bs}

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

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
