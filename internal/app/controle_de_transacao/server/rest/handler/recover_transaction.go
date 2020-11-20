package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/response"
	"net/http"
)

type RecoverTransaction struct {
	service service.RecoverTransaction
}

func NewRecoverTransaction(service service.RecoverTransaction) *RecoverTransaction {
	return &RecoverTransaction{
		service: service,
	}
}

func (h *RecoverTransaction) Handle(c echo.Context) error {
	context := c.Request().Context()
	transactionID := c.Param("transactionId")

	transaction, err := h.service.Execute(context, transactionID)
	if err != nil {
		return err
	}

	resp := new(response.Transaction).FromModelDomain(transaction)
	return c.JSON(http.StatusOK, resp)
}
