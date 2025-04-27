package seeders

import (
	"cashier-backend-go/internal/common/utils"
	"cashier-backend-go/internal/domain/auth/model"
	"log"

	"gorm.io/gorm"
)

func SeedAdminUser(DB *gorm.DB) {
	role := model.Role{
		Name:        "Admin",
		Description: "Administrator role with full access",
	}

	var existingRole model.Role
	if err := DB.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := DB.Create(&role).Error; err != nil {
				log.Fatalf("Error seeding role: %v", err)
			}
			log.Println("Admin role created successfully!")
		} else {
			log.Fatalf("Error checking for existing role: %v", err)
		}
	} else {
		log.Println("Admin role already exists!")
	}

	admin := model.User{
		Username: "admin",
		Password: "rahasia",
		RoleID:   existingRole.ID,
	}

	var existingUser model.User
	if err := DB.Where("role_id = ?", existingRole.ID).First(&existingUser).Error; err != nil {
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
