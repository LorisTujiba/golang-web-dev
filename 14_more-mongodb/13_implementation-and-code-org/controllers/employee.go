package controllers

import (
	"net/http"
	"github.com/LorisTujiba/golang-web-dev/14_more-mongodb/13_implementation-and-code-org/model"
	"database/sql"
	"html/template"
	"encoding/json"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Insert(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"insert.gohtml",nil)
}

func ShowJSON(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	emps,err := model.GetAllEmployees()
	if err!=nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(emps)
	if err!= nil{
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w,"json.gohtml",string(bs))
}

func Employees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	emps,err := model.GetAllEmployees()
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w,"index.gohtml",emps)

}