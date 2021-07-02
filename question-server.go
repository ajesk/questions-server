package main

import (
	"fmt"
	"log"
	"net/http"
)

func delog(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yolo")
}

func handleRequests() {
	fmt.Println("beginning server thing")
	http.HandleFunc("/", delog)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("starting up server")
	handleRequests()
}
