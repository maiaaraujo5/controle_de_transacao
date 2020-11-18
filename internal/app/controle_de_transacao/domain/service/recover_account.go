package service

import (
	"context"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type RecoverAccount interface {
	Execute(parentContext context.Context, accountID string) (*model.Account, error)
}

type recoverAccount struct {
	repository repository.Account
}

func NewRecoverAccount(repository repository.Account) RecoverAccount {
	return recoverAccount{
		repository: repository,
	}
}

func (r recoverAccount) Execute(parentContext context.Context, accountID string) (*model.Account, error) {

	account, err := r.repository.Find(parentContext, accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
