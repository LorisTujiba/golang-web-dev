package main

import (
	"encoding/json"
	"fmt"
)

/*
Unmarshall, transform json to be placed into go data structure
 */

type user struct{
	Name string
	Age int
	Address string
	Gender string
}

func main(){
	var uData user

	datas := `{
		"Name":"Joseph",
		"Age":22,
		"Address":"KH Syahdan 104A",
		"Gender":"Male"
	}` // supposed that this data are received from an API

	err := json.Unmarshal([]byte(datas), &uData) //unmarshal, remember save to variable. Because unmarshall
	if err!=nil{
		panic(err)
	}

	fmt.Println(uData)//print the uData
	fmt.Println(uData.Gender)//get certain data
}