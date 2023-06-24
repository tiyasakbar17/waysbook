package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Book").Find(&transactions).Error

	return transactions, err
}
