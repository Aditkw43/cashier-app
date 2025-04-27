package repository

import (
	"cashier-backend-go/internal/domain/transaction/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
	FindByID(id uint) (*model.Transaction, error)
	FindAll() ([]model.Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: DB}
}

func (r *transactionRepository) Create(transaction *model.Transaction) (*model.Transaction, error) {
	err := r.DB.Create(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *transactionRepository) FindByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.DB.Preload("TransactionDetails").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) FindAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.DB.Preload("TransactionDetails").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
