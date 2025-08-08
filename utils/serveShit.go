package utils

import (
	"encoding/json"
	"net/http"
)

func ServeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ServeError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}
