package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/response"
	"net/http"
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

	account, err := h.service.Execute(context, accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	resp := new(response.Account).FromModelDomain(account)

	return c.JSON(http.StatusOK, resp)
}
