package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type CreateTransaction interface {
	Execute(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type createTransaction struct {
	repository repository.Transaction
}

func NewCreateTransaction(repository repository.Transaction) CreateTransaction {
	return createTransaction{
		repository: repository,
	}
}

func (c createTransaction) Execute(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error) {

	log.Info("transaction created")

	transaction, err := c.repository.Save(parentContext, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
