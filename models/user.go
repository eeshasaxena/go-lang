package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId  `json:"id" bson:"_id"`
	Name string   `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string   `json:"password" bson:"password"`
}

type Post struct {
	Id bson.ObjectId  `json:"id" bson:"_id"`
	Caption string `json:"caption" bson:"caption"`
	ImageURL string `json:"image_url" bson:"image_url"`
	Time time.Time `json:"time" bson:"time"`
	Userid string `json:"userid" bson:"userid"` 
}

