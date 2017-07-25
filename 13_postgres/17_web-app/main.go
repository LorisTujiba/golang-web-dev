package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
	"net/http"
)

//we're going to display the employees data
//so, create a struct

type Employee struct{
	Id int
	Name string
	Score int
	Salary int
}

var db *sql.DB

func main(){
	var err error
	db, err = sql.Open("postgres","postgres://postgres:admin@localhost/company?sslmode=disable")
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	http.HandleFunc("/employees",employeeIndex)
	http.ListenAndServe(":8080",nil)

}

func employeeIndex(w http.ResponseWriter, r *http.Request){

	if r.Method != "GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}

	//select query
	rows, err := db.Query("Select * from employees")
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	eml := make([]Employee,0)

	for rows.Next(){
		em := Employee{}
		err:= rows.Scan(&em.Id,&em.Name,&em.Score,&em.Salary)//order matters
		if err !=nil{
			panic(err)
		}
		eml = append(eml,em)
	}

	//print out
	for _, em := range eml{
		fmt.Fprintf(w,"%d : %s\t,Score : %d, IDR %d\n", em.Id,em.Name,em.Score,em.Salary)
	}
}
