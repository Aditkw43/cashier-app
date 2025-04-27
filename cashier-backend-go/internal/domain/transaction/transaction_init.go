package transaction

import (
	"cashier-backend-go/internal/domain/transaction/controller"
	"cashier-backend-go/internal/domain/transaction/repository"
	"cashier-backend-go/internal/domain/transaction/routes"
	"cashier-backend-go/internal/domain/transaction/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTransaction(DB *gorm.DB, r *gin.Engine) {
	transactionRepo := repository.NewTransactionRepository(DB)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionController := controller.NewTransactionController(transactionService)

	routes.NewTransactionRoutes(transactionController).TransactionRoutes(r)
}
