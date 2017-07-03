package main

/*==============================================================================
Case
==============================================================================
Create a data structure to pass to a template which contains information about
restaurant's menu including Breakfast, Lunch, and Dinner items
*/

import (
	"html/template"
	"log"
	"os"
)

type food struct {
	Name  string
	Price float64
}

type meal struct {
	Time string
	F    []food
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	meals := []meal{
		meal{
			"Breakfast",
			[]food{
				food{
					"Nasi Goreng",
					15000,
				},
				food{
					"Bubur Ayam",
					15000,
				},
				food{
					"Nasi Uduk",
					15000,
				},
			},
		},
		meal{
			"Lunch",
			[]food{
				food{
					"Ayam Bakar",
					18000,
				},
				food{
					"Ayam Goreng",
					18000,
				},
			},
		},
		meal{
			"Dinner",
			[]food{
				food{
					"Mie Ayam",
					16000,
				},
				food{
					"Mie Goreng",
					18000,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, meals)
	if err != nil {
		log.Fatal(err)
	}

}
