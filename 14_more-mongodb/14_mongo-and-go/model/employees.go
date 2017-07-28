package model

import (
	"net/http"
	"errors"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//DB exported
var DB *mgo.Database
var Employees *mgo.Collection

func init(){
	s,err := mgo.Dial("mongodb://admin:admin@localhost:27017/company")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("company")
	Employees = DB.C("employees")
}

//Employee model
type Employee struct {
	ID     int
	Name   string
	Score  int
	Salary int
}

func GetAllEmployees()([]Employee,error){

	emps := []Employee{}
	err := Employees.Find(bson.M{}).All(&emps)
	if err != nil{
		return nil,err
	}

	return emps,nil
}

func GetCertainEmployee(r *http.Request) (Employee,error){
	emp := Employee{}
	name := r.FormValue("name")
	if name == "" {
		return emp, errors.New("400. Bad Request.")
	}
	err := Employees.Find(bson.M{"name": name}).One(&emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func PutEmployee(r *http.Request) (Employee, error) {
	// get form values
	emp := Employee{}
	emp.Name = r.FormValue("name")
	emp.Score,_ = strconv.Atoi(r.FormValue("score"))
	emp.Salary,_ = strconv.Atoi(r.FormValue("salary"))

	// validate form values
	if emp.Name == "" || emp.Score == 0 || emp.Salary == 0{
		return emp, errors.New("400. Bad request. All fields must be complete.")
	}

	// insert values
	var err error
	err = Employees.Insert(emp)
	if err != nil {
		return emp, errors.New("500. Internal Server Error." + err.Error())
	}
	return emp, nil
}

func UpdateEmployee(r *http.Request) (Employee, error) {
	// get form values
	emp := Employee{}
	emp.Name = r.FormValue("name")
	emp.Score,_ = strconv.Atoi(r.FormValue("score"))
	emp.Salary,_ = strconv.Atoi(r.FormValue("salary"))

	if emp.Name == "" || emp.Score == 0 || emp.Salary == 0{
		return emp, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// update values
	var err error
	err = Employees.Update(bson.M{"name": emp.Name}, &emp)
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func DeleteEmployee(r *http.Request) error {
	name := r.FormValue("name")
	if name == "" {
		return errors.New("400. Bad Request.")
	}

	err := Employees.Remove(bson.M{"name": name})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}