package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //< don't forget this
	"net/http"
)

/*

Query a single row!, didn't have to loop again

Arguments to the SQL function are referenced in the function body using the
syntax $n: $1 refers to the first argument, $2 to the second, and so on.
If an argument is of a composite type, then the dot notation, e.g.,
$1.name, can be used to access attributes of the argument.
The arguments can only be used as data values, not as
identifiers.source: postgres docs

"
	 Behind the scenes, db.QueryRow (and also db.Query() and db.Exec())
	 work by creating a new prepared statement on the database,
	 and subsequently execute that prepared statement using
	 the placeholder parameters provided. This means
	 that all three methods are safe from SQL
	 injection when used correctly .

	 From Wikipedia:
	 Prepared statements are resilient against SQL injection, because
	 parameter values, which are transmitted later using a different
	 protocol, need not be correctly escaped. If the original
	 statement template is not derived from external input,
	 injection cannot occur.

	 The placeholder parameter syntax differs depending on your database.
	 Postgres uses the $N notation, but MySQL, SQL Server and others
	 use the ? character as a placeholder.

" - Alex Edwards

Run thr application with this,
curl -i localhost:8080/employees/emp?id=1

*/

//Employee model
type Employee struct {
	ID     int
	Name   string
	Score  int
	Salary int
}

var db *sql.DB

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
	http.ListenAndServe(":8080", nil)

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
