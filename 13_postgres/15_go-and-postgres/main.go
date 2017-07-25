package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
)

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

	fmt.Println("Connected!")
}
