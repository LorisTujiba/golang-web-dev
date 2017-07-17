package models

//User exported, so other package can use
type User struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	ID     string `json:"id"`
}
