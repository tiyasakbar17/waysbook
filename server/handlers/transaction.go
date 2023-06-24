package handlers

import (
	"net/http"
	resultdto "waysbook/dto/result"
	transactiondto "waysbook/dto/transaction"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) CreateNewTransaction(c echo.Context) error {
	userId, _ := c.Param("userId")

	carts, err := h.CartRepository.FindCarts(userId)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transaction := models.Transaction{
		IdUser: strconv.Atoi(userId)
	}
	validation := validator.New()

	err = validation.Struct(transaction)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, err:= h.TransactionRepository.CreateTransaction(transaction, carts)

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: convertResponseTransaction(transaction),
	})
}

func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		Id:     u.Id,
	}
}
