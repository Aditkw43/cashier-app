package seeders

import (
	"gorm.io/gorm"
)

func InitSeed(db *gorm.DB) {
	SeedAdminUser(db)
}
