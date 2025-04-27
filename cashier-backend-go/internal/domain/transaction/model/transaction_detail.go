package model

import (
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionID uint    `gorm:"not null"`
	MenuID        uint    `gorm:"not null"`
	Quantity      int     `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	Total         float64 `gorm:"not null"`
}
