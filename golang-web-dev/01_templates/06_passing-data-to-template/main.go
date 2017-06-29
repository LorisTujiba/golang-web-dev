package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("index.gohtml"))
}

func main() {

	err := tpl.Execute(os.Stdout, 42)
	if err != nil {
		log.Fatal(err)
	}

}
