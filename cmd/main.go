package main

import (
	"net/http"

	"github.com/ImnotEdMateo/guestbook/db"
	"github.com/ImnotEdMateo/guestbook/routes"
	"github.com/gorilla/mux"
)

func main () {
  db.DBConnect()
  db.DBMigrate()
  
  r := mux.NewRouter()

  r.HandleFunc("/", routes.GetEntriesHandler).Methods("GET")
  r.HandleFunc("/", routes.PostEntryHandler).Methods("POST")
  r.HandleFunc("/entry/{id}", routes.GetEntryHandler).Methods("GET")

  http.ListenAndServe(":3000", r)
}
