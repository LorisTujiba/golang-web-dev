package model

import (
	"net/http"
	"errors"
	"database/sql"
	_ "github.com/lib/pq"
)

//DB exported
var DB *sql.DB

func init(){
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:admin@localhost/company?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

//Employee model
type Employee struct {
	ID     int
	Name   string
	Score  int
	Salary int
}

func GetAllEmployees()([]Employee,error){

	rs,err := DB.Query("SELECT * FROM employees")
	if err != nil{
		return nil,err
	}

	defer rs.Close()

	emps := make([]Employee,0)

	for rs.Next(){
		emp:= Employee{}
		rs.Scan(&emp.ID,&emp.Name,&emp.Score,&emp.Salary)
		if err != nil{
			return nil,err
		}
		emps = append(emps, emp)
	}

	if err = rs.Err(); err != nil {
		return nil, err
	}

	return emps,nil
}

func GetCertainEmployee(r *http.Request) (Employee,error){
	emp := Employee{}
	id := r.FormValue("id")

	if id == ""{
		return emp,errors.New("400. Bad Request")
	}

	rs,err := DB.Query("SELECT * FROM employees WHERE ID = $1")
	if err!=nil{
		return emp,err
	}

	err = rs.Scan(&emp.ID,&emp.Name,&emp.Score,&emp.Salary)
	if err!=nil{
		return emp,err
	}

	return emp,nil

}