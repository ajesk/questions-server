package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	http.HandleFunc("/ruok", healthCheck)
	http.HandleFunc("/poll", PollHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
