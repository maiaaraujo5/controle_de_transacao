package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/errors"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/request"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/server/rest/model/response"
	"net/http"
)

type CreateAccount struct {
	service  service.CreateAccount
	validate *validator.Validate
}

func NewCreateAccount(service service.CreateAccount, validate *validator.Validate) *CreateAccount {
	return &CreateAccount{
		service:  service,
		validate: validate,
	}
}

func (h *CreateAccount) Handle(c echo.Context) error {
	context := c.Request().Context()
	req, err := request.NewAccount(c)
	if err != nil {
		return errors.NewErrorResponse(c, http.StatusBadRequest, "bad request")
	}

	err = h.validate.Struct(req)
	if err != nil {
		return errors.NewErrorResponse(c, http.StatusBadRequest, "bad request")
	}

	account := &model.Account{
		DocumentNumber: req.DocumentNumber,
	}

	account, err = h.service.Execute(context, account)
	if err != nil {
		return errors.ToErrorResponse(c, err)
	}

	resp := new(response.Account).FromModelDomain(account)

	return c.JSON(http.StatusCreated, resp)
}
