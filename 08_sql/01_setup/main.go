package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //give it an alias _
	"io"
	"net/http"
)

/*
Install mysql community server and then get the go package
so later we can set the driver, run this

go get github.com/go-sql-driver/mysql

*/

func main() {

	//the connection string are username:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := sql.Open("mysql", "root:mypassword@tcp(localhost:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Successfully completed.")
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
