package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/request"
)

type CreateAccount struct {
	service service.CreateAccount
}

func NewCreateAccount(service service.CreateAccount) *CreateAccount {
	return &CreateAccount{
		service: service,
	}
}

func (h *CreateAccount) Handle(c echo.Context) error {
	context := c.Request().Context()

	req, err := request.NewAccount(c)
	if err != nil {
		return err
	}

	account := &model.Account{
		DocumentNumber: req.DocumentNumber,
	}

	account, err = h.service.Execute(context, account)
	if err != nil {
		return err
	}

	return nil
}
