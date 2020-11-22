package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/handler"
)

func Routes(createAccount *handler.CreateAccount, recoverAccount *handler.RecoverAccount,
	createTransaction *handler.CreateTransaction, recoverTransaction *handler.RecoverTransaction) *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())

	authenticateRoutes := e.Group("")
	authenticateRoutes.Use(middleware.JWT([]byte("key-segura")))

	e.POST("/accounts", createAccount.Handle)

	authenticateRoutes.GET("/accounts/:accountId", recoverAccount.Handle)
	authenticateRoutes.POST("/transactions", createTransaction.Handle)
	authenticateRoutes.GET("/transactions/:transactionId", recoverTransaction.Handle)
	return e
}
