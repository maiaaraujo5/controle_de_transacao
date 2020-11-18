package provider

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
)

type Account struct {
}

func NewAccount() repository.Account {
	return Account{}
}

func (a Account) Save(parentContext context.Context, account *model.Account) (*model.Account, error) {
	log.Info("saved")
	return nil, nil
}

func (a Account) Find(parentContext context.Context, accountID string) (*model.Account, error) {
	log.Info("found")
	return nil, nil
}
