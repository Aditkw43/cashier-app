package routes

import (
	"cashier-backend-go/internal/domain/transaction/controller"

	"github.com/gin-gonic/gin"
)

type TransactionRoutes struct {
	controller *controller.TransactionController
}

func NewTransactionRoutes(controller *controller.TransactionController) *TransactionRoutes {
	return &TransactionRoutes{controller: controller}
}

func (r *TransactionRoutes) TransactionRoutes(router *gin.Engine) {
	transactionGroup := router.Group("/transactions")
	{
		transactionGroup.POST("/", r.controller.CreateTransaction)
		transactionGroup.GET("/:id", r.controller.GetTransactionByID)
		transactionGroup.GET("/", r.controller.GetAllTransactions)
	}
}
