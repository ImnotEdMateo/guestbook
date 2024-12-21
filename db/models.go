package db

import (
  "time"

  "gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Name      string    `gorm:"default:Anonymous"` 
	Message   string    `gorm:"not null"`          
	Website   string
  CreatedAt time.Time 
}


