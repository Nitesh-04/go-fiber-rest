package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt time.Time
}