package main

/*==============================================================================
Case
==============================================================================
Using the data structure created in the previous folder, modify it to hold
menu information for an unlimited number of restaurants.
*/

import (
	"html/template"
	"os"
	"log"
)

type food struct{
	Name string
	Price float64
}

type meal struct{
	Time string
	F []food
}

type menu []meal

type restaurant struct{
	Name string
	Menu menu
}

type restaurants []restaurant

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){

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

	meals2 := []meal{
		meal{
			"Breakfast",
			[]food{
				food{
					"Nasi Campur",
					15000,
				},
				food{
					"Bubur Sosis",
					15000,
				},
				food{
					"Nasi Kucing",
					15000,
				},
			},
		},
		meal{
			"Lunch",
			[]food{
				food{
					"Ikan Bakar",
					18000,
				},
				food{
					"Ikan Goreng",
					18000,
				},
			},
		},
		meal{
			"Dinner",
			[]food{
				food{
					"Nasi Ayam",
					16000,
				},
				food{
					"Nasi Goreng",
					18000,
				},
			},
		},
	}

	restaurants := restaurants{
		restaurant{
			"A",
			meals2,
		},
		restaurant{
			"B",
			meals,
		},
	}

	err := tpl.Execute(os.Stdout,restaurants)
	if err != nil{
		log.Fatal(err)
	}

}
