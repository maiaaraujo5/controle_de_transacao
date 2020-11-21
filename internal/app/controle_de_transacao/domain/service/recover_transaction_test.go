package service

import (
	"context"
	"errors"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
	"time"
)

type RecoverTransactionSuite struct {
	suite.Suite
}

func TestNewRecoverTransactionSuite(t *testing.T) {
	suite.Run(t, new(RecoverTransactionSuite))
}

func (s *RecoverTransactionSuite) TestNewRecoverTransaction() {
	repo := new(mocks.Transaction)
	type args struct {
		repository repository.Transaction
	}
	tests := []struct {
		name string
		args args
		want RecoverTransaction
	}{
		{
			name: "should success build NewRecoverTransaction",
			args: args{
				repository: repo,
			},
			want: &recoverTransaction{
				repository: repo,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewRecoverTransaction(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRecoverTransaction() = %v, want %v", got, tt.want)
		})
	}
}

func (s *RecoverTransactionSuite) Test_recoverTransaction_Execute() {
	type fields struct {
		repository *mocks.Transaction
	}
	type args struct {
		parentContext context.Context
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Transaction
		wantErr bool
		mock    func(repository *mocks.Transaction)
	}{
		{
			name: "should success recover an transaction",
			fields: fields{
				repository: new(mocks.Transaction),
			},
			args: args{
				parentContext: context.Background(),
				transactionID: mock.Anything,
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          300,
				EventDate:       time.Date(2020, time.November, 20, 20, 07, 35, 0, time.UTC),
			},
			wantErr: false,
			mock: func(repository *mocks.Transaction) {
				repository.On("Find", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          300,
					EventDate:       time.Date(2020, time.November, 20, 20, 07, 35, 0, time.UTC),
				}, nil).Once()
			},
		},
		{
			name: "should return error when repository return error",
			fields: fields{
				repository: new(mocks.Transaction),
			},
			args: args{
				parentContext: context.Background(),
				transactionID: mock.Anything,
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.Transaction) {
				repository.On("Find", mock.Anything, mock.Anything).Return(nil, errors.New("error to recover a transaction")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.repository)

			r := &recoverTransaction{
				repository: tt.fields.repository,
			}

			got, err := r.Execute(tt.args.parentContext, tt.args.transactionID)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v")
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Execute() = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
