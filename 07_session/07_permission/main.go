package main

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

/*
Create a permission based on role

so at the struct, we add new field called role
*/

var tpl *template.Template

type user struct {
	Username string
	Password []byte
	Role     string
}

var userDatas = map[string]user{}
var association = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	//pre-defined user

	pass, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	user := user{
		"admin",
		pass,
		"Admin",
	}
	userDatas[user.Username] = user

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/transaction", transaction)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func transaction(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := getUser(w, r)

	if user.Role != "Admin" {
		http.Error(w, "Only admin can be authorized", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "transaction.gohtml", nil)

}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")

	//delete from association
	delete(association, c.Value)

	//remove cookie
	c = &http.Cookie{
		Name:   "session",
		MaxAge: -1,
		Value:  "",
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	//validate first, if already logged in, redirect
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		//get the value
		un := r.FormValue("username")
		up := r.FormValue("password")

		//check if the password match
		if bcrypt.CompareHashAndPassword(userDatas[un].Password, []byte(up)) != nil {
			http.Error(w, "Doesn't Match", http.StatusUnauthorized)
			return
		}

		//else, then put the data into session
		SID := uuid.NewV4()
		sess := &http.Cookie{
			Name:  "session",
			Value: SID.String(),
		}
		http.SetCookie(w, sess)

		//put it in the association
		association[sess.Value] = un

		//then redirect
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	user := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func home(w http.ResponseWriter, r *http.Request) {

	user := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "home.gohtml", user)

}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) { // if user already logged in , redirect
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		role := r.FormValue("role")

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

		userDatas[username] = user{username, bs, role}

		//redirect
		http.Redirect(w, r, "/home", http.StatusSeeOther)
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
		http.SetCookie(w, cookie)
	}

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
