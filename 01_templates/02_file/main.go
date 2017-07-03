package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	//Save the HTML into another file

	name := "Loris Tujiba"
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name + //passing the data
		`</h1>
		</body>
		</html>
	`)

	nf, err := os.Create("index.html") //create a file with the name of index.html
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
