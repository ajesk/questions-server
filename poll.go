package main

import (
	"fmt"
	"net/http"
)

type Poll struct {
	Id     string `json:"id"`
	AltId  string `json:"altId"`
	Code   string `json:"code"`
	Status string `json:"status"`
	Link   string `json:"link"`
	Name   string `json:"name"`
}

func createPoll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create poll hit")
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
