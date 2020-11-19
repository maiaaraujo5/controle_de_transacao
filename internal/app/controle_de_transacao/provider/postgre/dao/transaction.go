package dao

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type Transaction struct {
	db *pg.DB
}

func NewTransaction(db *pg.DB) repository.Transaction {
	return Transaction{db: db}
}

func (t Transaction) Save(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	_, err := t.db.Model(transaction).Insert()
	if err != nil {
		return nil, err
	}
	log.Info("saved")
	return transaction, nil
}

func (t Transaction) Find(parentContext context.Context, transactionID string) (*model.Transaction, error) {
	log.Info("found")
	return nil, nil
}
