package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
	"html/template"
	"net/http"
	"strconv"
)

/*
	/update?id=[desired id]
*/

//Employee model
type Employee struct {
	ID     int
	Name   string
	Score  int
	Salary int
}

var db *sql.DB
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:admin@localhost/company?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/update", update)
	http.HandleFunc("/employees", employees)
	http.HandleFunc("/", insert)
	http.ListenAndServe(":8080", nil)
}

func update(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	rs := db.QueryRow("SELECT * FROM employees WHERE id = $1", id)
	emp := Employee{}
	rs.Scan(&emp.ID, &emp.Name, &emp.Score, &emp.Salary)

	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		score := r.FormValue("score")
		salary := r.FormValue("salary")
		db.Exec("UPDATE employees SET name=$1,score=$2,salary=$3 WHERE  id = $4", name, score, salary, id)

		w.Header().Set("Location", "/employees")
		w.WriteHeader(http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "update.gohtml", emp)

}

func insert(w http.ResponseWriter, r *http.Request) {

	var err error
	if r.Method == http.MethodPost {
		emp := Employee{}
		emp.Name = r.FormValue("name")
		emp.Score, err = strconv.Atoi(r.FormValue("score"))
		if err != nil {
			panic(err)
		}
		emp.Salary, err = strconv.Atoi(r.FormValue("salary"))
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("INSERT INTO employees (name,score,salary) VALUES($1,$2,$3)", emp.Name, emp.Score, emp.Salary)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "insert.gohtml", nil)
}

func employees(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//select query
	rows, err := db.Query("Select * from employees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	eml := make([]Employee, 0)

	for rows.Next() {
		em := Employee{}
		err := rows.Scan(&em.ID, &em.Name, &em.Score, &em.Salary) //order matters
		if err != nil {
			panic(err)
		}
		eml = append(eml, em)
	}

	//print out
	for _, em := range eml {
		fmt.Fprintf(w, `%d : %s\t,Score : %d, IDR %d \n`, em.ID, em.Name, em.Score, em.Salary)
	}
}
