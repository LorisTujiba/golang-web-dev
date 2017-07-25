package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
)

//we're going to display the employees data
//so, create a struct

type Employee struct{
	Id int
	Name string
	Score int
	Salary int
}

func main(){
	db, err := sql.Open("postgres","postgres://postgres:admin@localhost/company?sslmode=disable")
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
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
		fmt.Printf("%d : %s\t,Score : %d, IDR %d\n", em.Id,em.Name,em.Score,em.Salary)
	}

}
