package routes

import (
	"waysbook/handlers"
	mysql "waysbook/pkg/database"
	"waysbook/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(TransactionRepository)

	e.GET("/transaction", h.CreateNewTransaction)
}
