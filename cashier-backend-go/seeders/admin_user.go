package seeders

import (
	"cashier-backend-go/internal/common/utils"
	"cashier-backend-go/internal/domain/auth/model"
	"log"

	"gorm.io/gorm"
)

func SeedAdminUser(DB *gorm.DB) {
	admin := model.User{
		Username: "admin",
		Password: "rahasia",
		Role:     "admin",
	}

	var existingUser model.User
	if err := DB.Where("role = ?", "admin").First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPassword, err := utils.HashPassword(admin.Password)
			if err != nil {
				log.Fatalf("Error hashing password: %v", err)
			}
			admin.Password = hashedPassword

			if err := DB.Create(&admin).Error; err != nil {
				log.Fatalf("Error seeding admin user: %v", err)
			}

			log.Println("Admin user created successfully!")
		} else {
			log.Fatalf("Error checking for existing admin user: %v", err)
		}
	} else {
		log.Println("Admin user already exists!")
	}
}
