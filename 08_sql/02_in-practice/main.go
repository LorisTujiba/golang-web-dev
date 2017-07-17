package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //give it an alias _
	"io"
	"net/http"
)

/*
CRUD
*/

var db *sql.DB
var err error

func main() {

	//the connection string are username:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "root:mypassword@tcp(localhost:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/friends", friends)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)

}

func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM  pets where name = 'Josh'`)
	check(err)

	row, err := stmt.Exec() //execute statement and resulting a result
	check(err)

	n, err := row.RowsAffected() //get how many row(s) affected
	check(err)

	fmt.Fprintln(w, "Inserted into table pets, rows affected : ", n)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`Update pets set name='Josh',breed='Alaskan Malamute' where name = 'howard'`)
	check(err)

	row, err := stmt.Exec() //execute statement and resulting a result
	check(err)

	n, err := row.RowsAffected() //get how many row(s) affected
	check(err)

	fmt.Fprintln(w, "Inserted into table pets, rows affected : ", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO pets (name,breed) values('howard','German Shepperd')`)
	check(err)

	row, err := stmt.Exec() //execute statement and resulting a result
	check(err)

	n, err := row.RowsAffected() //get how many row(s) affected
	check(err)

	fmt.Fprintln(w, "Inserted into table pets, rows affected : ", n)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE pets(name varchar(20),breed varchar(20));`) //pointer to a statement and an error
	check(err)

	row, err := stmt.Exec() //execute statement and resulting a result
	check(err)

	n, err := row.RowsAffected() //get how many row(s) affected
	check(err)

	fmt.Fprintln(w, "CREATED Table pets, rows affected : ", n)
}

func friends(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM friends;`)
	check(err)

	var s, name, age string
	s = "Retrieved Records\n"

	for rows.Next() {
		err = rows.Scan(&name, &age)
		check(err)
		s += name + " : " + age + "\n"
	}
	fmt.Fprintln(w, s)

	rows, err2 := db.Query(`select * from countries`)
	check(err2)

	var cName string

	s = ""

	for rows.Next() {
		err = rows.Scan(&cName)
		check(err)

		s += cName + "\n"
	}

	fmt.Fprintln(w, s)

	rows, err3 := db.Query(`select * from pets`)
	check(err3)

	var pName, breed string

	s = ""

	for rows.Next() {
		err = rows.Scan(&pName, &breed)
		check(err)

		s += pName + " : " + breed + "\n"
	}

	fmt.Fprintln(w, s)

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
