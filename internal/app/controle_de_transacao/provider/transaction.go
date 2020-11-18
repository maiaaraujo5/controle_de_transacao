package provider

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type Transaction struct {
}

func NewTransaction() repository.Transaction {
	return Transaction{}
}

func (t Transaction) Save(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	log.Info("saved")
	return nil, nil
}

func (t Transaction) Find(parentContext context.Context, transactionID string) (*model.Transaction, error) {
	log.Info("found")
	return nil, nil
}
