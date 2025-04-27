package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string `gorm:"type:text"`
}
