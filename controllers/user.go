package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/MeganViga/GoWithMongoDB/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type UserController struct{
	Client *mongo.Client
}

func NewUsersController(client *mongo.Client)*UserController{
	return &UserController{
		Client: client,
	}
}

func (uc *UserController)GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !primitive.IsValidObjectID(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	u := models.User{}
	if err := uc.Client.Database("testdb").Collection("users").FindOne(context.Background(),bson.M{"_id":oid}).Decode(&u); err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)

}

func  (uc *UserController)CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// body, _ := ioutil.ReadAll(r.Body)
	u := models.User{}
	// json.Unmarshal(body, &u)
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = primitive.NewObjectID()
	uc.Client.Database("testdb").Collection("users").InsertOne(context.Background(),u)
	uj, err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type","applications/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

func (uc *UserController)DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !primitive.IsValidObjectID(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid, _ := primitive.ObjectIDFromHex(id)
	result, err := uc.Client.Database("testdb").Collection("users").DeleteOne(context.Background(),bson.M{"_id":oid})
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Println(result)

}