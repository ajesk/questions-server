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

type Interest struct {
	_id      string             `json: "_id" bson:"_id", "omitempty"`
	AltId    string             `json: "altId" bson:"altId"`
	Question primitive.ObjectID `json: "pollId" bson:"pollId", "omitempty"`
}

var interestCollection = "interest"

func toInterest(jsonString []byte) (Interest, error) {
	var interest Interest
	err := json.Unmarshal(jsonString, &interest)
	if err != nil {
		return Interest{}, err
	}
	return interest, nil
}

func CreateInterest(w http.ResponseWriter, r *http.Request) {
	pollId, _ := primitive.ObjectIDFromHex(mux.Vars(r)["pollId"])
	questionId, _ := primitive.ObjectIDFromHex(mux.Vars(r)["questionId"])

	pollExists := PollExists(w, pollId)
	if !pollExists {
		log.Println("poll does not exist aborting")
		return
	}

	questionExists := QuestionExists(w, questionId)
	if !questionExists {
		log.Println("questions does not exist aborting")
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	interest, err := toInterest(body)
	if err != nil {
		log.Println("error caught while parsing json body aborting", err)
		return
	}
	interest.Question = questionId
	fmt.Fprintf(w, insertInterest(interest))
}

func insertInterest(interest Interest) string {

	res, err := GetCollection(interestCollection).InsertOne(context.Background(), interest)
	if err != nil {
		log.Fatalln("error occurred while creating interest", err)
	}

	b, _ := json.Marshal(res.InsertedID)
	return string(b)
}
