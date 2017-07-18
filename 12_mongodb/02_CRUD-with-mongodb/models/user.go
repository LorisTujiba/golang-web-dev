package models

import "gopkg.in/mgo.v2/bson"

//User is exported so that controller can access
type User struct {
	Username string        `json:"username" bson:"name"`
	Gender   string        `json:"gender" bson:"gender"`
	Age      int           `json:"age" bson:"age"`
	ID       bson.ObjectId `json:"id" bson:"_id"` //change this to bson

	//ID was string before
}
