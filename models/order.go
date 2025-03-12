package models

import "time"

type Order struct {
	ID        uint `json:"id" gorm:"primary_key"`
	ProductID int `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    int  `json:"user_id"`
	User      User  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}