package dao

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
	DBModel "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/provider/postgre/model"
)

type Account struct {
	db *pg.DB
}

func NewAccount(db *pg.DB) repository.Account {
	return Account{db: db}
}

func (a Account) Save(parentContext context.Context, account *model.Account) (*model.Account, error) {

	accountDB := new(DBModel.Account).FromDomainModel(account)

	_, err := a.db.WithContext(parentContext).Model(accountDB).Returning("id").Insert()
	if err != nil {
		return nil, err
	}

	log.Info("saved")
	return accountDB.ToDomainModel(), nil
}

func (a Account) Find(parentContext context.Context, accountID string) (*model.Account, error) {

	accountDB := new(DBModel.Account)
	err := a.db.WithContext(parentContext).Model(accountDB).Where("id = ?0", accountID).Select()
	if err != nil {
		return nil, err
	}

	log.Info("found")
	return accountDB.ToDomainModel(), nil
}
