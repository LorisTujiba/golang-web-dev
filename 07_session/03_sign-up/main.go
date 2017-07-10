package main

import (
	"github.com/LorisTujiba/gotraining/src/github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

var tpl *template.Template

type user struct {
	Username string
	Password string
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

		userDatas[username] = user{username, password}

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
