package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*"))
}

type user struct {
	Name string
	Age  int
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml",42)
	if err != nil {
		log.Fatal(err)
	}

}
