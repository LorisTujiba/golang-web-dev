package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

/*=============================================
Session
=============================================
Store information to server. Client gonna send
server unique id. this unique id is called
SID, session id. and there will be a
storage area where the session id
have a association with user id.

create session that store a uuid using cookie!
we're creating uuid using satori's package

go get

run go get github.com/satori/go.uuid
*/

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session") //get a cookie named session
	//if there is no cookie named session, create one
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:   "session",
			Value:  id.String(),
			Secure: true, // run only if the connection is secure, only https.
			// If you want to test in localhost, please change to false or
			// comment the secure
			HttpOnly: true, //so this cookie cannot be accessed with javascript
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
