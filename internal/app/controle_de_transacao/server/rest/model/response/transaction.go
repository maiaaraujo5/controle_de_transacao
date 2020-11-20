package response

import (
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"time"
)

type Transaction struct {
	ID              int       `json:"id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float32   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func (t *Transaction) FromModelDomain(transaction *model.Transaction) *Transaction {
	t.ID = transaction.ID
	t.AccountID = transaction.AccountID
	t.OperationTypeID = transaction.OperationTypeID
	t.Amount = transaction.Amount
	t.EventDate = transaction.EventDate
	return t
}
