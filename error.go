package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Detail  string `json: "detail"`
}

func errorResponse(status int, w http.ResponseWriter, err ErrorResponse) {
	w.WriteHeader(status)
	errorJson, _ := json.Marshal(err)
	fmt.Fprintf(w, string(errorJson))
}

func BadRequestResponse(w http.ResponseWriter, err error, detail string) {
	w.WriteHeader(http.StatusBadRequest)
	response := ErrorResponse{err.Error(), detail}
	errorResponse(http.StatusNotFound, w, response)
}

func NotFoundResponse(w http.ResponseWriter, err error, detail string) {
	w.WriteHeader(http.StatusNotFound)
	response := ErrorResponse{err.Error(), detail}
	errorResponse(http.StatusNotFound, w, response)
}
