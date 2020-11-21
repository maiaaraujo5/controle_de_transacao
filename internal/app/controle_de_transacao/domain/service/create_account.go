package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type CreateAccount interface {
	Execute(parentContext context.Context, account *model.Account) (*model.Account, error)
}

type createAccount struct {
	repository repository.Account
}

func NewCreateAccount(repository repository.Account) CreateAccount {
	return &createAccount{
		repository: repository,
	}
}

func (c createAccount) Execute(parentContext context.Context, account *model.Account) (*model.Account, error) {

	account, err := c.repository.Save(parentContext, account)
	if err != nil {
		return nil, err
	}

	log.Info("created account")
	return account, nil
}
