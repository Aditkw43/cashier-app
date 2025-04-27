package controller

import (
	"cashier-backend-go/internal/common/request"
	"cashier-backend-go/internal/domain/transaction/model"
	"cashier-backend-go/internal/domain/transaction/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var input request.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transactionDetails := make([]model.TransactionDetail, len(input.TransactionDetails))
	for i, detail := range input.TransactionDetails {
		transactionDetails[i] = model.TransactionDetail{
			MenuID:   detail.MenuID,
			Quantity: detail.Quantity,
			Price:    detail.Price,
			Total:    detail.Total,
		}
	}

	transaction, err := c.transactionService.CreateTransaction(input.CustomerID, transactionDetails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

func (c *TransactionController) GetTransactionByID(ctx *gin.Context) {
	id := ctx.Param("id")
	transactionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := c.transactionService.GetTransactionByID(uint(transactionID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (c *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := c.transactionService.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
