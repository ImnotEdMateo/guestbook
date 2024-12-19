package db

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Name    string `gorm:"default:anonymous"` 
	Message string `gorm:"not null"`          
	Website string
	Email   string
}
