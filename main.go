package main

import (
	"os"
	"log"
  "net/http"
	"strings"

  "github.com/ImnotEdMateo/guestbook/db"
  "github.com/ImnotEdMateo/guestbook/routes"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

func main () {
	port := os.Getenv("GUESTBOOK_PORT")
	if port == "" {
		port = "3000"
	}

	origins := os.Getenv("GUESTBOOK_ALLOWED_ORIGINS")
	if origins == "" {
		origins = "http://localhost"
	}

	allowedOrigins := strings.Split(origins, ",")

  db.DBConnect()
  db.DBMigrate()

	log.Printf("Corriendo servidor en http://0.0.0.0:%s", port)

  r := mux.NewRouter()
  r.HandleFunc("/", routes.GetEntriesHandler).Methods("GET")
  r.HandleFunc("/", routes.PostEntryHandler).Methods("POST")
  r.HandleFunc("/entry/{id}", routes.GetEntryHandler).Methods("GET")


  handler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)
  
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
