package service

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model/operations_types"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
	"math"
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

	if !operations_types.IsValidOperationType(transaction.OperationTypeID) {
		return nil, errors.New("the operation type is invalid")
	}

	if transaction.OperationTypeID != operations_types.PAYMENT {
		transaction.Amount = math.Copysign(transaction.Amount, -1)
	}

	t, err := c.repository.Save(parentContext, transaction)
	if err != nil {
		return nil, err
	}

	log.Info("transaction created")
	return t, nil
}
