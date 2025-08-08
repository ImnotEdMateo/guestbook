package server

import (
	"net/http"

	"github.com/ImnotEdMateo/guestbook/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewServer(allowedOrigins []string) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetEntriesHandler).Methods("GET")
	r.HandleFunc("/", handlers.PostEntryHandler).Methods("POST")
	r.HandleFunc("/entry/{id}", handlers.GetEntryHandler).Methods("GET")

	return cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)
}
