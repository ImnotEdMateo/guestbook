package main

import (
	"log"
	"net/http"

	"github.com/ImnotEdMateo/guestbook/config"
	"github.com/ImnotEdMateo/guestbook/db"
	"github.com/ImnotEdMateo/guestbook/server"
)

func main() {
	cfg := config.Load()

	db.DBConnect()
	db.DBMigrate()

	log.Printf("Corriendo servidor en http://0.0.0.0:%s", cfg.Port)

	srv := server.NewServer(cfg.AllowedOrigins)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, srv))
}
