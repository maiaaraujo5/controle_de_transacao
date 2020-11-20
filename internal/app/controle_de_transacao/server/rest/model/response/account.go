package response

import "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"

type Account struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (a *Account) FromModelDomain(account *model.Account) *Account {
	a.AccountID = account.AccountID
	a.DocumentNumber = account.DocumentNumber
	return a
}
