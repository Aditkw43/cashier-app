package service

import (
	"cashier-backend-go/internal/domain/transaction/model"
	"cashier-backend-go/internal/domain/transaction/repository"
	"errors"
)

type TransactionService interface {
	CreateTransaction(customerID uint, details []model.TransactionDetail) (*model.Transaction, error)
	GetTransactionByID(id uint) (*model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) CreateTransaction(customerID uint, details []model.TransactionDetail) (*model.Transaction, error) {
	if customerID == 0 || len(details) == 0 {
		return nil, errors.New("invalid transaction data")
	}

	transaction := &model.Transaction{
		CustomerID:         customerID,
		TotalAmount:        calculateTotalAmount(details),
		TransactionDetails: details,
	}

	transaction, err := s.transactionRepo.Create(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	return s.transactionRepo.FindByID(id)
}

func (s *transactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.FindAll()
}

func calculateTotalAmount(details []model.TransactionDetail) float64 {
	var total float64
	for _, detail := range details {
		total += detail.Total
	}
	return total
}
