package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	resp := Health{Status: "up"}
	json.NewEncoder(w).Encode(resp)
}

func handleRequests() {
	fmt.Println("beginning server thing")
	http.HandleFunc("/ruok", healthCheck)
	http.HandleFunc("/poll", PollHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("starting up server")
	handleRequests()
}
