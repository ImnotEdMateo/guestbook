package db

import (
  "log"
   "os"
  
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
  host := os.Getenv("DB_HOST")
  user := os.Getenv("DB_USER") 
  password := os.Getenv("DB_PASSWORD") 
  dbname := os.Getenv("DB_NAME") 
  port := os.Getenv("DB_PORT")
  dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
    
  var error error
  DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if error != nil {
    log.Fatal(error)
  } else {
    log.Println("Se ha conectado exitosamente a PostgreSQL")
  }
}

func DBMigrate() {
  if err := DB.AutoMigrate(&Entry{}); err != nil {
    log.Fatalf("No se pudo migrar la base de datos: %v", err)
  }
  log.Println("Migración completada con éxito")
}
