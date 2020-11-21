package service

import (
	"bou.ke/monkey"
	"context"
	"errors"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model/operations_types"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"time"
)

type CreateTransactionSuite struct {
	suite.Suite
}

func TestCreateTransactionSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionSuite))
}

func (s *CreateTransactionSuite) SetupSuite() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, time.November, 20, 20, 07, 35, 0, time.UTC)
	})
}

func (s *CreateTransactionSuite) TestNewCreateTransaction() {
	transaction := new(mocks.Transaction)
	account := new(mocks.Account)

	type args struct {
		transactionRepository repository.Transaction
		accountRepository     repository.Account
	}
	tests := []struct {
		name string
		args args
		want CreateTransaction
	}{
		{
			name: "should success build NewCreateTransaction",
			args: args{
				transactionRepository: transaction,
				accountRepository:     account,
			},
			want: &createTransaction{
				transactionRepository: transaction,
				accountRepository:     account,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewCreateTransaction(tt.args.transactionRepository, tt.args.accountRepository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateTransaction() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateTransactionSuite) Test_createTransaction_Execute() {
	type fields struct {
		transaction *mocks.Transaction
		account     *mocks.Account
	}
	type args struct {
		parentContext context.Context
		transaction   *model.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Transaction
		wantErr bool
		mock    func(transaction *mocks.Transaction, account *mocks.Account)
	}{
		{
			name: "should success save a new transaction with operationTypeId cash purchase ",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          100,
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: operations_types.CASH_PUCRCHASE,
				Amount:          -100,
				EventDate:       time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -100,
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
		{
			name: "should success save a new transaction with operationTypeId installment_purchase ",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 2,
					Amount:          50,
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: operations_types.INSTALLMENT_PURCHASE,
				Amount:          -50,
				EventDate:       time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 2,
					Amount:          -50,
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
		{
			name: "should success save a new transaction with operationTypeId withdraw",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 3,
					Amount:          100,
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: operations_types.WITHDRAW,
				Amount:          -100,
				EventDate:       time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {

				account.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 3,
					Amount:          -100,
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
		{
			name: "should success save a new transaction with operationTypeId payment",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          100,
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: operations_types.PAYMENT,
				Amount:          100,
				EventDate:       time.Now(),
			},
			wantErr: false,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {

				account.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          100,
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
		{
			name: "should return error when transaction repository return error",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 4,
					Amount:          100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {

				account.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()

				transaction.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save transaction")).Once()
			},
		},
		{
			name: "should return error when operation type is invalid",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 5,
					Amount:          100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Find", mock.Anything, mock.Anything).Maybe()
				transaction.On("Save", mock.Anything, mock.Anything).Maybe()
			},
		},
		{
			name: "should return error when find account repository return error",
			fields: fields{
				transaction: new(mocks.Transaction),
				account:     new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				transaction: &model.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          100,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Find", mock.Anything, mock.Anything).Return(nil, errors.New("error to find account")).Once()
				transaction.On("Save", mock.Anything, mock.Anything).Maybe()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.transaction, tt.fields.account)

			c := createTransaction{
				accountRepository:     tt.fields.account,
				transactionRepository: tt.fields.transaction,
			}
			got, err := c.Execute(tt.args.parentContext, tt.args.transaction)

			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v")
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Execute() = %v, want %v", got, tt.want)

			tt.fields.account.AssertExpectations(s.T())
			tt.fields.transaction.AssertExpectations(s.T())
		})
	}
}
