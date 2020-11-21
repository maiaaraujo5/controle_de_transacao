package repository

import (
	"context"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
)

type Account interface {
	Save(parentContext context.Context, account *model.Account) (*model.Account, error)
	Find(parentContext context.Context, accountID string) (*model.Account, error)
	FindByDocumentNumber(parentContext context.Context, documentNumber string) (*model.Account, error)
}
