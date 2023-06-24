package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction, carts []models.Cart) (int, error)
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

func (r *repository) CreateTransaction(transaction models.Transaction, carts []models.Cart) (int, error) {
	err = r.db.Create(&transaction).Error
	if err != nil {
		return 0, err
	}
	transactionBooks := make([]models.TransactionBook, len(carts))

	for i, cart := range carts {
		transactionBooks[i] = models.TransactionBook{
			IdTransaction: transaction.Id,
			IdBook: cart.IdBook
		}
	}
	err = r.db.Create(&transactionBooks).Error

	return transactions.Id, err
}