package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model/operations_types"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/errors"
	"math"
	"strconv"
	"time"
)

type CreateTransaction interface {
	Execute(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type createTransaction struct {
	transactionRepository repository.Transaction
	accountRepository     repository.Account
}

func NewCreateTransaction(transactionRepository repository.Transaction, accountRepository repository.Account) CreateTransaction {
	return &createTransaction{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

func (c createTransaction) Execute(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error) {

	if !operations_types.IsValidOperationType(transaction.OperationTypeID) {
		return nil, errors.BadRequest("the operation type is invalid")
	}

	_, err := c.accountRepository.Find(parentContext, strconv.Itoa(transaction.AccountID))
	if err != nil {
		return nil, err
	}

	if transaction.OperationTypeID != operations_types.PAYMENT {
		transaction.Amount = math.Copysign(transaction.Amount, -1)
	}
	transaction.EventDate = time.Now()
	t, err := c.transactionRepository.Save(parentContext, transaction)
	if err != nil {
		return nil, err
	}

	log.Info("transaction created")
	return t, nil
}
