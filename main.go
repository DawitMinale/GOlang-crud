package main

import (
	"log"
	"mongo-golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)

}

func getSession() *mgo.Session {
	uri := "mongodb+srv://deva13:1310@cluster0.2t8nvqw.mongodb.net/go-web-dev-db"

	s, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return s
}
