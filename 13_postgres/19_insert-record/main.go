package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
	"html/template"
	"net/http"
	"strconv"
)


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
	tpl = template.Must(template.ParseFiles("insert.gohtml"))
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

	http.HandleFunc("/employees", employeeIndex)
	http.HandleFunc("/employees/emp", employeeEmp)
	http.HandleFunc("/", insert)
	http.ListenAndServe(":8080", nil)

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

	tpl.Execute(w, nil)
}

func employeeEmp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(404), http.StatusBadRequest)
		return
	}

	//select query
	row := db.QueryRow("Select * from employees where id = $1", id)

	em := Employee{}

	err := row.Scan(&em.ID, &em.Name, &em.Score, &em.Salary) //order matters
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%d : %s\t,Score : %d, IDR %d\n", em.ID, em.Name, em.Score, em.Salary)
}

func employeeIndex(w http.ResponseWriter, r *http.Request) {

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
		fmt.Fprintf(w, "%d : %s\t,Score : %d, IDR %d\n", em.ID, em.Name, em.Score, em.Salary)
	}
}
