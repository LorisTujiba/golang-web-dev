package main

/*================================
| - Pipeline
================================
With pipeline we can pass an
output value as an input
for another function.
*/

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"square":   sqrt,
	"division": division,
}

func sqrt(input float64) float64 {
	return input * input
}

func division(input float64) float64 {
	return input / input
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now()) //get today's date
	if err != nil {
		log.Fatal(err)
	}

}
