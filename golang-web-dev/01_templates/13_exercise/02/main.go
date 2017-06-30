package main

/*==============================================================================
Case
==============================================================================
Create a data structure to pass to a template which contains information about
California hotels including Name, Address, City, Zip, Region region can be:
Southern, Central, Northern

can hold an unlimited number of hotels
*/

import (
	"html/template"
	"os"
	"log"
)

type Hotel struct{
	Name string
	Address string
	City string
	Zip string
	Region string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main(){

	datas := []Hotel{
		Hotel{"A",
			"A Street",
			"A City",
			"11530",
			"Central"},
		Hotel{"Central",
			  "A",
			  "Central City",
			  "12320",
			  "Central"},
		Hotel{"B",
			  "B",
			  "B",
			  "13530",
			  "Central"},
	}

	err := tpl.Execute(os.Stdout,datas)
	if err != nil{
		log.Fatal(err)
	}

}
