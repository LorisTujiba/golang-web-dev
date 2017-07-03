package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() { // func init run only once when the program start
	//so we are ensuring that this templates only run once
	tpl = template.Must(template.ParseGlob("templates/*")) //Get all of the file that located in templates/

	/*-------------------------------------------------------------
	func Must

	func Must(t *Template, err error) *Template
	Must is a helper that wraps a call to a function returning
	(*Template, error) and panics if the error is non-nil.
	It is intended for use in variable initializations such as

	var t = template.Must(template.New("name").Parse("text"))

	If we open the Must code, it's doing an error checking,
	so we can skip the error checking and instead using
	the Must.
	--------------------------------------------------------------*/
}

func main() {

	err := tpl.Execute(os.Stdout, nil) //print the first only
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "products.gohtml", nil) //get file with the name products
	if err != nil {
		log.Fatal(err)
	}

}
