package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", rootUrl)
	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func rootUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := APIStatus{Status: "OK", Message: "Connected to API"}
	json.NewEncoder(w).Encode(response)
}
