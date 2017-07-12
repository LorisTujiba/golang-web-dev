package main

import (
"github.com/LorisTujiba/gotraining/src/github.com/satori/go.uuid"
"html/template"
"net/http"
"golang.org/x/crypto/bcrypt"
	"time"
	"fmt"
)

/*======================================================
Expire
======================================================
What if we want to make a session that will expire in
a certain time ?. We can use this code below to make
that happen.
*/

var tpl *template.Template

type user struct {
	Username string
	Password []byte
	Role string
}

//create session with field to track their last activity
//time
type session struct{
	un string
	lastActivity time.Time
}

var userDatas = map[string]user{}
var association = map[string]session{}//change from string to session, we're no longer just using the username
var cleanTime time.Time

//Create a constant variable as the time length
const sessionLength int = 30


func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	//pre-defined user

	//We're initiating that the cleanTime is when the program first start, act as the pivot point
	cleanTime = time.Now()

	pass,err := bcrypt.GenerateFromPassword([]byte("admin"),bcrypt.MinCost)
	if err!=nil{
		panic(err)
	}

	user := user{
		"admin",
		pass,
		"Admin",
	}
	userDatas[user.Username] = user

}

func logout(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	c,_ := r.Cookie("session")

	//delete from association
	delete(association,c.Value)

	//remove cookie
	c = &http.Cookie{
		Name:"session",
		MaxAge:-1,
		Value:"",
	}

	http.SetCookie(w,c)

	//clean up dbsessions
	if time.Now().Sub(cleanTime) > (time.Second *30){//if the time of the program started and the clean time differ 30 secs, clean the session
		//clean every 30 sec
		go cleanSessions()
	}

	http.Redirect(w,r,"/",http.StatusSeeOther)
	return
}

func login(w http.ResponseWriter,r *http.Request){
	//validate first, if already logged in, redirect
	if alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost{
		//get the value
		un := r.FormValue("username")
		up := r.FormValue("password")

		//check if the password match
		if bcrypt.CompareHashAndPassword(userDatas[un].Password,[]byte(up))!=nil{
			http.Error(w,"Doesn't Match",http.StatusUnauthorized)
			return
		}

		//else, then put the data into session
		SID := uuid.NewV4()
		sess := &http.Cookie{
			Name:"session",
			MaxAge:sessionLength,
			Value:SID.String(),
		}
		http.SetCookie(w,sess)

		//put it in the association
		association[sess.Value] = session{un,time.Now()}

		//then redirect
		http.Redirect(w,r,"/home",http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w,"login.gohtml",nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w,r) { // if user already logged in , redirect
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
			MaxAge: sessionLength,
			Value: SID.String(),
		}
		http.SetCookie(w, cookie)

		association[cookie.Value] = session{username,time.Now()}

		//use bcrypt
		bs,err :=bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		userDatas[username] = user{username, bs,role}

		//redirect
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func cleanSessions(){
	showSessions()//shows the list of the sessions
	for k,v := range association{
		if time.Now().Sub(v.lastActivity) > 3000 || v.un == ""{// if the time of program start and the user last act differ more than 30 secs, delete the assoc
			delete(association,k)
		}
	}
	cleanTime = time.Now()// refresh the clean time to now
	fmt.Println("After Clean")
	showSessions()//shows the list of the sessions
}

func showSessions(){
	for k,v := range association{
		fmt.Println(k,v.un)
	}
}

//To refresh the last activity time, we place the code
//at these 2 funcs, because all of our routes uses this
//funcs

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	cookie, err := req.Cookie("session")
	if err != nil {//if there's no session, create the session
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}

	//if there's already a session
	// if the user exists already, get user
	var u user
	if ud, ok := association[cookie.Value]; ok {
		//change to now
		ud.lastActivity = time.Now()
		//replace
		association[cookie.Value] = ud
		u = userDatas[ud.un]
	}

	//reset the lastact back to 30 sec
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	cookie, err := req.Cookie("session") //get the cookie
	if err != nil {//if there's already a cookie, redirect
		return false
	}
	ud,ok := association[cookie.Value]
	if ok{
		ud.lastActivity = time.Now()
		association[cookie.Value] = ud
	}

	_, okay := userDatas[ud.un] //check if the cookie contain the username, then use ok idiom

	cookie.MaxAge = sessionLength
	http.SetCookie(w,cookie)

	return okay
}


//Add show session to track

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/transaction", transaction)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login",login)
	http.HandleFunc("/logout",logout)
	http.ListenAndServe(":8080", nil)
}

func transaction(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	user := getUser(w,r)

	if user.Role != "Admin"{
		http.Error(w,"Only admin can be authorized",http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w,"transaction.gohtml",nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	user := getUser(w, r)
	showSessions()
	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func home(w http.ResponseWriter, r *http.Request) {

	user := getUser(w, r)
	if !alreadyLoggedIn(w,r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(w, "home.gohtml", user)

}
