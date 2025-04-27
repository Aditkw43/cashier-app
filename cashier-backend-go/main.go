package main

import (
	"cashier-backend-go/config"
	"cashier-backend-go/internal/domain/auth"
	authModel "cashier-backend-go/internal/domain/auth/model"
	"cashier-backend-go/internal/domain/menu"
	menuModel "cashier-backend-go/internal/domain/menu/model"
	"cashier-backend-go/internal/domain/transaction"
	transactionModel "cashier-backend-go/internal/domain/transaction/model"
	"cashier-backend-go/seeders"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	if err := config.DB.AutoMigrate(
		&authModel.User{},
		&authModel.Role{},
		&menuModel.Menu{},
		&transactionModel.Transaction{},
		&transactionModel.TransactionDetail{},
	); err != nil {
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
	transaction.InitTransaction(config.DB, r)
}
