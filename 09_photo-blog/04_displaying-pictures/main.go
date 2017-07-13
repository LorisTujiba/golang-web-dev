package main

import (
	"net/http"
	"html/template"
	"time"
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"//give it an alias _
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
)

/*
Displaying, change the home.gohtml
and strip the public prefix
that's it
 */

type user struct{
	Username string
	Password []byte
	Gender string
	DateOfBirth time.Time
}

type session struct{
	Name string
}

var uSession = map[string]user{}//save the logged in users
var dbSession = map[string]session{}//save the sessions
var db *sql.DB
var err error
var tpl *template.Template
const sessionLength = 30

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){

	db,err = sql.Open("mysql","root:mypassword@tcp(localhost:3306)/photoblog?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.Handle("/public/",http.StripPrefix("/public",http.FileServer(http.Dir("./public"))))//strip the public, localhost:.../public, to serve the files manually
	http.HandleFunc("/",index)
	http.HandleFunc("/signup",signUp)
	http.HandleFunc("/logout",logout)
	http.HandleFunc("/home",home)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func getUserData(w http.ResponseWriter,r *http.Request) *http.Cookie{
	ud,err:= r.Cookie("userData")
	var sid = uuid.NewV4()
	if err!=nil{
		ud = &http.Cookie{
			Name:"userData",
			Value:sid.String(),
		}
		http.SetCookie(w,ud)
	}

	return ud
}

func home(w http.ResponseWriter,r *http.Request){
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	c := getUserData(w,r)

	var images = make([]string,3)
	images = append(images, "loris.jpg")
	images = append(images, "dog.jpg")

	if !strings.Contains(c.Value,"loris.jpg") ||!strings.Contains(c.Value,"dog.jpg") {
		for _,image :=range images{

				c.Value += "|" + image
		}
	}

	http.SetCookie(w,c)

	pass := strings.Split(c.Value,"|")


	if r.Method == http.MethodPost{
		mf,fileheader,err := r.FormFile("nf")
		check(err)
		defer mf.Close()

		//get the file extension
		ext := strings.Split(fileheader.Filename,".")[1]//1 because splitted, do not take the 0 , cuz thats the name
		//[nameofthefile][jpg]

		//create sha1 hash
		h := sha1.New()
		io.Copy(h,mf)//copy (desti,source)
		fname:= fmt.Sprintf("%x",h.Sum(nil)) + "." + ext //hexadecimal value %x

		//create new file
		wd,err := os.Getwd() //get working directory
		check(err)
		path := filepath.Join(wd,"public","pics",fname)//save at public/pics/[location]

		nf,err := os.Create(path)
		check(err)
		defer nf.Close()

		//copy, the file, but
		//cuz of the hash  has travers the file, we reset it to zero
		mf.Seek(0,0)
		io.Copy(nf,mf)

		if !strings.Contains(c.Value,fname){
			c.Value += "|"+fname
		}
		http.SetCookie(w,c)

		http.Redirect(w,r,"/home",http.StatusSeeOther)
		return
	}


	tpl.ExecuteTemplate(w,"home.gohtml",pass)
}

func logout(w http.ResponseWriter,r *http.Request) {
	if !alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	c,err := r.Cookie("loggedInUser")
	check(err)

	c.MaxAge = -1
	delete(dbSession,c.Value)

	http.Redirect(w,r,"/",http.StatusSeeOther)
	return

}

func signUp(w http.ResponseWriter,r *http.Request){
	if alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/",http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost{
		username := r.FormValue("username")
		password := r.FormValue("password")
		gender := r.FormValue("gender")
		dob := r.FormValue("dob")

		//check if the username is taken
		res,err := db.Query("SELECT username FROM users WHERE username = ?;",username)
		check(err)

		if res.Next(){
			http.Error(w,"Username's already taken",http.StatusForbidden)
			return
		}

		//store user's data to db
		stmt,err := db.Prepare(`INSERT INTO users (username,password,gender,dob) values(?,?,?,?)`)
		check(err)

		//bcrypt thr password
		p,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.MinCost)
		check(err)

		row,err := stmt.Exec(username,p,gender,dob)
		check(err)

		n,err := row.RowsAffected()//get how many row(s) affected
		check(err)

		fmt.Fprintln(w,"Success! rows affected : ",n)
		return
	}
	tpl.ExecuteTemplate(w,"signup.gohtml",nil)
}

func index(w http.ResponseWriter,r *http.Request){
	if alreadyLoggedIn(w,r){
		http.Redirect(w,r,"/home",http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost{
		username := r.FormValue("username")
		password := r.FormValue("password")

		res,err := db.Query("SELECT * FROM users WHERE username=?",username)
		check(err)

		var u user

		for res.Next(){
			res.Scan(&u.Username,&u.Password,&u.Gender,&u.DateOfBirth)
		}
		if bcrypt.CompareHashAndPassword(u.Password,[]byte(password))!=nil{
			http.Error(w,"Password Doesn't Match",http.StatusUnauthorized)
			return
		}

		sid := uuid.NewV4()
		c := &http.Cookie{
			Name:"loggedInUser",
			Value:sid.String(),
			MaxAge:sessionLength,
		}
		http.SetCookie(w,c)

		dbSession[c.Value] = session{username}
		uSession[u.Username] = u
		http.Redirect(w,r,"/home",http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request)bool{

	c,err := r.Cookie("loggedInUser")
	if err !=nil{
		return false
	}

	un := dbSession[c.Value]

	c.MaxAge = sessionLength
	http.SetCookie(w,c)

	_,ok := uSession[un.Name]

	return ok
}

func check(err error){
	if err!=nil{
		panic(err.Error())
	}
}
