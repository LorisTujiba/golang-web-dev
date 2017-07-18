package models

//User , user model, exported
type User struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
