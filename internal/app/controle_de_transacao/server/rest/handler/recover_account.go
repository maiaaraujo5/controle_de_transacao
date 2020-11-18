package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
)

type RecoverAccount struct {
	service service.RecoverAccount
}

func NewRecoverAccount(service service.RecoverAccount) *RecoverAccount {
	return &RecoverAccount{
		service: service,
	}
}

func (h *RecoverAccount) Handle(c echo.Context) error {
	context := c.Request().Context()

	accountID := c.Param("accountId")

	_, err := h.service.Execute(context, accountID)
	if err != nil {
		return err
	}

	return nil
}
