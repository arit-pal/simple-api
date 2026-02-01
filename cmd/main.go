package main

import (
	"log"
	"net/http"
	"simple-api/internal/server"
)

func main() {
	handler := server.New()

	log.Println("Server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
