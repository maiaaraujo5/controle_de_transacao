package model

import "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"

type Account struct {
	tableName      struct{} `pg:"accounts"`
	ID             int      `pg:"id,pk"`
	DocumentNumber string   `pg:"document_number"`
}

func (a *Account) FromDomainModel(account *model.Account) *Account {
	a.ID = account.AccountID
	a.DocumentNumber = account.DocumentNumber

	return a
}

func (a *Account) ToDomainModel() *model.Account {
	return &model.Account{
		AccountID:      a.ID,
		DocumentNumber: a.DocumentNumber,
	}
}
