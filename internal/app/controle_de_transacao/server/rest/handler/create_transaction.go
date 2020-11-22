package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	restErrors "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/errors"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/request"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/response"
	"net/http"
)

type CreateTransaction struct {
	service  service.CreateTransaction
	validate *validator.Validate
}

func NewCreateTransaction(service service.CreateTransaction, validate *validator.Validate) *CreateTransaction {
	return &CreateTransaction{
		service:  service,
		validate: validate,
	}
}

func (h *CreateTransaction) Handle(c echo.Context) error {
	context := c.Request().Context()

	req, err := request.NewTransaction(c)
	if err != nil {
		return restErrors.NewErrorResponse(c, http.StatusBadRequest, "bad request")
	}

	err = h.validate.Struct(req)
	if err != nil {
		return restErrors.NewErrorResponse(c, http.StatusBadRequest, "bad request")
	}

	transaction := &model.Transaction{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
	}

	transaction, err = h.service.Execute(context, transaction)
	if err != nil {
		return restErrors.ToErrorResponse(c, err)
	}

	resp := new(response.Transaction).FromModelDomain(transaction)

	return c.JSON(http.StatusCreated, resp)
}
