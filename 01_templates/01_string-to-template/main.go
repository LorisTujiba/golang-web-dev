package main

import "fmt"

func main() {
	/*=======================================================================
	Template
	=======================================================================
	Creating template for the site, to do that we need to set the template
	using the html. Template commonly used as a center point while
	delivering data. Very much like blade in laravel.
	*/

	name := "Loris Tujiba"

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(template)
}
