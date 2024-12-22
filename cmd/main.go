package main

import (
  "net/http"

  "github.com/ImnotEdMateo/guestbook/db"
  "github.com/ImnotEdMateo/guestbook/routes"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

func main () {
  db.DBConnect()
  db.DBMigrate()
  
  r := mux.NewRouter()
  r.HandleFunc("/", routes.GetEntriesHandler).Methods("GET")
  r.HandleFunc("/", routes.PostEntryHandler).Methods("POST")
  r.HandleFunc("/entry/{id}", routes.GetEntryHandler).Methods("GET")


  handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://edmateo.site"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(r)
  
  http.ListenAndServe(":3000", handler)
}
