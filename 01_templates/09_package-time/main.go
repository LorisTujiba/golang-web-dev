package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"formatDate":    mdy,
	"formatKitchen": kit,
}

func mdy(input time.Time) string {
	return input.Format("11-03-1995") //return the inputted time into this format
}

func kit(input time.Time) string {
	return input.Format(time.Kitchen) //package time predefined function
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
