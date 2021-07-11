package main

import (
	"context"
	"fmt"
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

func createPoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create poll hit")
	collection := GetCollection("test")
	res, err := collection.InsertOne(context.Background(), Poll{"42", "43", "44", "pending", "dot com", "yolo"})
	if err != nil {
		fmt.Println("fam plz", err)
	}

	fmt.Println("did the thing", res)
}

func endPoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("end poll hit")
}

func getPoll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get poll hit")
}

func PollHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPoll(w, r)
	case "POST":
		createPoll(w, r)
	case "DELETE":
		endPoll(w, r)
	}
}
