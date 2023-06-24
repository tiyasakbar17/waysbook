package handlers

import (
	"net/http"
	resultdto "waysbook/dto/result"
	transactiondto "waysbook/dto/transaction"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateNewTransaction(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseTransaction(transactions),
	})
}

func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Id:     u.Id,
		IdUser: u.IdUser,
		IdBook: u.IdBook,
	}
}
