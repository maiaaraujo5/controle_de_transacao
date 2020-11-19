package dao

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type Account struct {
	db *pg.DB
}

func NewAccount(db *pg.DB) repository.Account {
	return Account{db: db}
}

func (a Account) Save(parentContext context.Context, account *model.Account) (*model.Account, error) {
	_, err := a.db.Model(account).Insert()
	if err != nil {
		return nil, err
	}
	log.Info("saved")
	return account, nil
}

func (a Account) Find(parentContext context.Context, accountID string) (*model.Account, error) {
	log.Info("found")
	return nil, nil
}
