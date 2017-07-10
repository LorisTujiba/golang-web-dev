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
	http.HandleFunc("/", setSession)
	http.HandleFunc("/getuserdata", getSession)
	http.ListenAndServe(":8080", nil)
}

func setSession(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("session")
	id := uuid.NewV4()
	var u user
	if err != nil {
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	if uname, ok := association[cookie.Value]; ok {
		u = userDatas[uname]
	}

	if r.Method == http.MethodPost {
		userName := r.FormValue("name")
		userPass := r.FormValue("password")
		association[cookie.Value] = userName
		u = user{userName, userPass}
		userDatas[userName] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func getSession(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil { //if no cookie, redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	uname, ok := association[c.Value] // get the username
	if !ok {                          // if the desired username doesn't have cookie, redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := userDatas[uname]
	tpl.ExecuteTemplate(w, "home.gohtml", u)

}
