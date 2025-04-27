package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID         uint                `gorm:"not null"`
	TotalAmount        float64             `gorm:"not null"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
