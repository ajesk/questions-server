package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Interest struct {
	_id      string             `json: "_id" bson:"_id", "omitempty"`
	AltId    string             `json: "altId" bson:"altId"`
	Question primitive.ObjectID `json: "pollId" bson:"pollId"`
}

var collectionName = "interest"

func CreateInterest(w http.ResponseWriter, r *http.Request) {
	//body, _ := ioutil.ReadAll(r.Body)

	pollId, _ := primitive.ObjectIDFromHex(mux.Vars(r)["pollId"])
	questionId, _ := primitive.ObjectIDFromHex(mux.Vars(r)["questionId"])
	altId, _ := mux.Vars(r)["altId"]

	// poll exists
	pollExists := PollExists(w, pollId)
	if !pollExists {
		log.Println("poll does not exist aborting")
		return
	}
	// question exists

	fmt.Fprintf(w, insertInterest(questionId, altId))
}

func insertInterest(questionId primitive.ObjectID, altId string) string {
	interest := Interest{}
	interest.AltId = altId
	interest.Question = questionId

	res, err := GetCollection(collectionName).InsertOne(context.Background(), interest)
	if err != nil {
		log.Fatalln("error occurred while creating interest", err)
	}

	b, _ := json.Marshal(res.InsertedID)
	return string(b)
}
