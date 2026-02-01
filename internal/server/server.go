package server

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/hello", HelloHandler)

	return LoggingMiddleware(mux)
}
