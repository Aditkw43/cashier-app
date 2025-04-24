package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Category    string  `gorm:"type:varchar(255);not null"`
}
