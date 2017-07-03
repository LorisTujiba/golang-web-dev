package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parsing from files, so we don't have to write the HTML code here. Which is a pain.
	//So instead we can parse the html file and manipulate it here.

	tpl, err := template.ParseFiles("index.gohtml") //add the file to the container
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html") //creating file
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close() // defer, wait till the main ready to exit

	// If the template is only one, can use the execute
	err = tpl.Execute(os.Stdout, nil) // print to terminal
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(nf, nil) //print to nf, the new file that we create before
	if err != nil {
		log.Fatal(err)
	}

	//If there are more than one, use Execute Template to specifies which file you want to execute
	tpl, err = tpl.ParseFiles("home.gohtml") // add one more

	err = tpl.ExecuteTemplate(os.Stdout, "home.gohtml", nil) // print to terminal
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil) // print to terminal
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, nil) // Only print the first index of the container
	if err != nil {
		log.Fatal(err)
	}

}
