package models

import (
	// "fmt"
	// "github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	// "net/http"
)

// User is the model for a user account

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Gender   string `json:"gender" bson:"gender"`
	Age      int   `json:"age" bson:"age"`
}
