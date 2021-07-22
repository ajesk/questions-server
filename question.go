package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	_id    string             `json: "_id" bson:"_id", "omitempty"`
	AltId  string             `json: "altId" bson:"altId"`
	Poll   primitive.ObjectID `json: "pollId" bson:"pollId"`
	Status string             `json: "status" bson:"status", "omitempty"`
	Text   string             `json: "text" bson:"text", "omitempty"`
}

var collectionName = "question"

func toQuestion(jsonString string) Question {
	var question Question
	err := json.Unmarshal([]byte(jsonString), &question)
	if err != nil {
		log.Fatalln(err)
	}
	return question
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["pollId"])

	log.Println(id)
	pollExists := PollExists(w, id)
	if !pollExists {
		log.Println("poll does not exist aborting")
		return
	}

	question := toQuestion(string(body))
	question.Status = "open"
	question.Poll = id

	res, err := GetCollection(collectionName).InsertOne(context.Background(), question)
	if err != nil {
		log.Fatalln("error occurred while creating question", err)
	}

	b, _ := json.Marshal(res.InsertedID)
	fmt.Fprintf(w, string(b))
}

func InterestQuestion(w http.ResponseWriter, r *http.Request) {
	log.Println("interest question hit")
}
