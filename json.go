package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
		w.WriteHeader(500)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	if statusCode >= 500 {
		log.Println("Server Error:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errorResponse{Error: message})
}
