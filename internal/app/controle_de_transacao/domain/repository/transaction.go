package repository

import (
	"context"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
)

type Transaction interface {
	Save(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error)
	Find(parentContext context.Context, transactionID string) (*model.Transaction, error)
}
