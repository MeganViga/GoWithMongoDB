package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MeganViga/GoWithMongoDB/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
r := httprouter.New()
uc := controllers.NewUsersController(getSession())
r.GET("/user/:id",uc.GetUser)
r.POST("/user",uc.CreateUser)
r.DELETE("/user/:id",uc.DeleteUser)
log.Fatal(http.ListenAndServe(":9091",r))
}
func getSession()*mongo.Client{
	uri := "mongodb://user:mongopass@localhost:27017"
	client, err := mongo.Connect(context.TODO(),options.Client().ApplyURI(uri))
	if err != nil{
		panic(err)
	}
	return client
}