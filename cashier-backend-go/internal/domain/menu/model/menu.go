package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null"`
	Category    string  `gorm:"not null"`
	IsActive    bool    `gorm:"default:true"`
}
