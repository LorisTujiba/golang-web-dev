package main

import (
	"fmt"
	"net/http"
)

/*=============================================
Cookie
=============================================
SetCookie accept response writer and a Cookie

A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an HTTP response or the Cookie header of an HTTP request.

See http://tools.ietf.org/html/rfc6265 for details.

type Cookie struct {
        Name  string
        Value string

        Path       string    // optional
        Domain     string    // optional
        Expires    time.Time // optional
        RawExpires string    // for reading cookies only

        // MaxAge=0 means no 'Max-Age' attribute specified.
        // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
        // MaxAge>0 means Max-Age attribute present and given in seconds
        MaxAge   int
        Secure   bool
        HttpOnly bool
        Raw      string
        Unparsed []string // Raw text of unparsed attribute-value pairs
}
*/

func main() {
	http.HandleFunc("/", setCook)
	http.HandleFunc("/read", readCook)
	http.ListenAndServe(":8080", nil)
}

func setCook(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: "Joseph",
	})
	fmt.Println(w, "Cookie is written, chek the browser, dev tools / application / cookies")
}

func readCook(w http.ResponseWriter, r *http.Request) {
	co, err := r.Cookie("username")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintln(w, "here is the cookie , ", co)
}
