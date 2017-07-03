package main

/*=================================
Predefined function
================================
A Function that you can use in a
template.
*/

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

	//Check the html
	//index, range, if

	xs := []string{"zero", "one", "two", "three"}

	data := struct {
		Words []string
		LName string
	}{
		xs,
		"Loris",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", xs)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "field.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

	a := user{
		"John",
		23,
	}

	b := user{
		"Doe",
		17,
	}

	c := user{
		"",
		0,
	}

	users := []user{a, b, c}

	err = tpl.ExecuteTemplate(os.Stdout, "if.gohtml", users)
	if err != nil {
		log.Fatal(err)
	}

}
