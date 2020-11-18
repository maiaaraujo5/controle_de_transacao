package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	fxHandler "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/fx/module/rest/handler"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/handler"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fxHandler.AccountModule(),
		fxHandler.TransactionModule(),
		fx.Provide(
			Routes,
		),
	)
}

func Routes(createAccount *handler.CreateAccount, recoverAccount *handler.RecoverAccount,
	createTransaction *handler.CreateTransaction) *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/accounts", createAccount.Handle)
	e.GET("/accounts", recoverAccount.Handle, middleware.JWT([]byte("key-segura")))
	e.POST("/transactions", createTransaction.Handle, middleware.JWT([]byte("key-segura")))
	return e
}
