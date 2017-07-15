package main

import (
	"encoding/json"
	"fmt"
)

/*
IN AN HTML FILE

Create an array to hold objects in javascript.

Each object should store a "Code" as a number and "Descrip" as a string.

The array should hold all of the objects. Create objects for the following HTTP status codes:

    StatusOK                   = 200
    StatusMovedPermanently  = 301
    StatusFound             = 302
    StatusSeeOther          = 303
    StatusTemporaryRedirect = 307
    StatusBadRequest                   = 400
    StatusUnauthorized                 = 401
    StatusPaymentRequired              = 402
    StatusForbidden                    = 403
    StatusNotFound                     = 404
    StatusMethodNotAllowed             = 405
    StatusTeapot                       = 418
    StatusInternalServerError           = 500

Stringify that data to JSON

Display that JSON in the browser window

IN A GO FILE

Create a variable with the identifier "rcvd" of type string.

Store a raw string literal of the JSON created in the previous step as the value of the variable "rcvd".

Unmarshal "rcvd" into a data structure with the identifier "data"

Use a for range loop to iterate through "data" displaying the results to the terminal

*/

type Datas struct {
	Code int `json:"Code"`
	Descrip string `json:"Descrip"`
}

func main(){
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`


	var data []Datas

	err := json.Unmarshal([]byte(rcvd),&data)
	if err!=nil{
		panic(err)
	}

	for _,v := range data{
		fmt.Println(v.Code, "-", v.Descrip)
	}


}