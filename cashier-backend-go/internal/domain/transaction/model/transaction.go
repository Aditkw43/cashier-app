package model

import (
	"cashier-backend-go/internal/domain/auth/model"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID         uint                `gorm:"not null"`
	User               model.User          `gorm:"foreignKey:CustomerID"`
	TotalAmount        float64             `gorm:"not null"`
	TransactionDetails []TransactionDetail `gorm:"foreignKey:TransactionID"`
}
