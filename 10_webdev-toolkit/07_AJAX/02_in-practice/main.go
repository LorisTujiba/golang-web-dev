package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/checkName", checkName)
	http.ListenAndServe(":8080", nil)
}

func checkName(w http.ResponseWriter, r *http.Request) {

	sampleUsers := map[string]bool{
		"test@example.com": true,
		"jame@bond.com":    true,
		"moneyp@uk.gov":    true,
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(w, sampleUsers[sbs])

}

func foo(w http.ResponseWriter, r *http.Request) {

	tpl.Execute(w, nil)

}
