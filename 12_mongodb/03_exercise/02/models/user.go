package models

import (
	"encoding/json"
	"fmt"
	"os"
)

//User , user model, exported
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

//StoreUsers , store the data to the model
func StoreUsers(m map[string]User) {
	f, err := os.Create("data")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(m)
}

//LoadUsers , load datas
func LoadUsers() map[string]User {
	m := make(map[string]User)

	f, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return m
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	return m
}
