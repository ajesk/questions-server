package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Poll struct {
	Id     string `bson:"id"`
	AltId  string `bson:"altId"`
	Code   string `bson:"code"`
	Status string `bson:"status"`
	Link   string `bson:"link"`
	Name   string `bson:"name"`
}

var collection = "test"

func CreatePoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create poll hit")
	collection := GetCollection(collection)
	res, err := collection.InsertOne(context.Background(), Poll{"42", "43", "44", "pending", "dot com", "yolo"})
	if err != nil {
		fmt.Println("error occurred while creating poll", err)
	}

	b, _ := json.Marshal(res.InsertedID)
	fmt.Fprintf(w, string(b))
}

func EndPoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("end poll hit")
}

func GetPoll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get poll hit")
	log.Print(r)
}
