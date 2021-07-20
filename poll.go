package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Poll struct {
	_id    string `json: "_id" bson:"_id", "omitempty"`
	AltId  string `json: "altId" bson:"altId"`
	Code   string `json: "code" bson:"code", "omitempty"`
	Status string `json: "status" bson:"status", "omitempty"`
	Link   string `json:"link" bson:"link", "omitempty"`
	Name   string `json: "name" bson:"name"`
}

var collection = "poll"

func toPoll(jsonString string) Poll {
	var poll Poll
	err := json.Unmarshal([]byte(jsonString), &poll)
	if err != nil {
		log.Fatalln(err)
	}
	return poll
}

func CreatePoll(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	poll := toPoll(string(body))
	poll.Status = "active"
	collection := GetCollection(collection)

	res, err := collection.InsertOne(context.Background(), poll)
	if err != nil {
		log.Fatalln("error occurred while creating poll", err)
	}

	b, _ := json.Marshal(res.InsertedID)
	fmt.Fprintf(w, string(b))
}

func EndPoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("end poll hit")
}

func GetPoll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get poll hit")

	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var poll Poll
	err := GetCollection(collection).FindOne(context.Background(), bson.M{"_id": id}).Decode(&poll)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response := ErrorResponse{err.Error()}
		errorRet, _ := json.Marshal(response)
		fmt.Fprintf(w, string(errorRet))
		return
	}
	result, err := json.Marshal(poll)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
	fmt.Fprintf(w, string(result))
}

func PollExists(w http.ResponseWriter, id primitive.ObjectID) bool {
	var poll Poll
	err := GetCollection(collection).FindOne(context.Background(), bson.M{"_id": id}).Decode(&poll)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			response := ErrorResponse{err.Error()}
			errorRet, _ := json.Marshal(response)
			fmt.Fprintf(w, string(errorRet))
			return false
		}
	}
	return true
}
