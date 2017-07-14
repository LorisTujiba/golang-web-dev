package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

/*================================================
JSON, Javascript Object Notation
================================================
Nowadays used by all programming language to
transfer data back and forth. get the
package JSON!
*/
type person struct {
	Fname string
	Lname string
	Items []string
}

func main(){

	//marshall-un marshall, save it in a variable, stores the result in the value pointed to by v
	//encode-decode, write it somewhere

	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshl)
	http.HandleFunc("/encode", encod)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w,"You are at foo")
}

func marshl(w http.ResponseWriter, req *http.Request) {

	//set data to be parsed into JSON
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}

	//do marshal, remember, save to variable. If you want to directly write the data, use json
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encod(w http.ResponseWriter, req *http.Request) {

	//set data to be parsed into JSON
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}

	//do encode, automatically write
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
