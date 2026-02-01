package server

import (
	"net/http"
	"simple-api/internal/response"
	"time"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{
		"time": time.Now().Format(time.RFC3339),
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		response.WriteJSONError(w, http.StatusBadRequest, "name query parameter is required")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Hello, " + name,
	})
}
