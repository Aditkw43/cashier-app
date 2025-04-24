package main

import (
	"cashier-backend-go/config"
	"cashier-backend-go/internal/auth"
	"cashier-backend-go/internal/auth/model"
	"cashier-backend-go/internal/menu"
	menuModel "cashier-backend-go/internal/menu/model"
	"cashier-backend-go/seeders"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	if err := config.DB.AutoMigrate(&model.User{}, &menuModel.Menu{}); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	seeders.InitSeed(config.DB)

	r := gin.Default()
	InitDomain(r)

	// Run server
	r.Run(":8080")
}

func InitDomain(r *gin.Engine) {
	auth.InitAuth(config.DB, r)
	menu.InitMenu(config.DB, r)
}
