package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"time": time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Momo"
	}

	response := map[string]string{
		"message": "Hello, " + name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
