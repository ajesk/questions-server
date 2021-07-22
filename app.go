package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Health struct {
	Status string `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	resp := Health{Status: "up"}
	json.NewEncoder(w).Encode(resp)
}

func handleRequests() {
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	fmt.Println("beginning server thing")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ruok", healthCheck)

	router.HandleFunc("/poll", CreatePoll).Methods("POST")
	router.HandleFunc("/poll", CreatePoll).Methods("DELETE")
	router.HandleFunc("/poll/{id}", GetPoll).Methods("GET")
	router.HandleFunc("/poll/{pollId}/question", CreateQuestion).Methods("POST")
	router.HandleFunc("/poll/{pollId}/question/{questionId}", CreateInterest).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
