package __string_to_template

import "fmt"

func main(){
	name := "Loris Tujiba"

	template :=`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<title>Hello World!</title>
	</head>
	<body>
	<h1>`+name+`</h1>
	</body>
	</html>
	`
	fmt.Println(template)
}
