package model

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"time"
)

type Transaction struct {
	tableName       struct{}  `pg:"transactions"`
	ID              int       `pg:"id"`
	AccountID       int       `pg:"account_id"`
	OperationTypeID int       `pg:"operation_type_id"`
	Amount          float32   `pg:"amount"`
	EventDate       time.Time `pg:"event_date"`
}

func (t *Transaction) FromModelDomain(transaction *model.Transaction) *Transaction {
	t.AccountID = transaction.AccountID
	t.Amount = transaction.Amount
	t.OperationTypeID = transaction.OperationTypeID
	return t
}

func (t *Transaction) ToModelDomain() *model.Transaction {
	return &model.Transaction{
		AccountID:       t.AccountID,
		OperationTypeID: t.OperationTypeID,
		Amount:          t.Amount,
	}
}
