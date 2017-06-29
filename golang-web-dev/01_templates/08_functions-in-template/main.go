package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

//use funcmap to register functions
//uc is what the func is going to be called in the template
//ft is the function we declared

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {

	names := []string{"Maria", "Carissa", "Lidwina", "Tessa", "Claudia"}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}

}
