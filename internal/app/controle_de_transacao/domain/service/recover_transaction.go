package service

import (
	"context"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type RecoverTransaction interface {
	Execute(parentContext context.Context, transactionID string) (*model.Transaction, error)
}

type recoverTransaction struct {
	repository repository.Transaction
}

func NewRecoverTransaction(repository repository.Transaction) RecoverTransaction {
	return &recoverTransaction{repository: repository}
}

func (r *recoverTransaction) Execute(parentContext context.Context, transactionID string) (*model.Transaction, error) {

	transaction, err := r.repository.Find(parentContext, transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
