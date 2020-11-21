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
)

type RecoverAccountSuite struct {
	suite.Suite
}

func TestRecoverAccountSuite(t *testing.T) {
	suite.Run(t, new(RecoverAccountSuite))
}

func (s *RecoverAccountSuite) TestNewRecoverAccount() {

	repo := new(mocks.Account)

	type args struct {
		repository repository.Account
	}
	tests := []struct {
		name string
		args args
		want RecoverAccount
	}{
		{
			name: "should success build NewRecoverAccount",
			args: args{
				repository: repo,
			},
			want: &recoverAccount{
				repository: repo,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewRecoverAccount(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRecoverAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *RecoverAccountSuite) Test_recoverAccount_Execute() {
	type fields struct {
		repository *mocks.Account
	}
	type args struct {
		parentContext context.Context
		accountID     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
		mock    func(repository *mocks.Account)
	}{
		{
			name: "should success recover an account",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				accountID:     mock.Anything,
			},
			want: &model.Account{
				AccountID:      10,
				DocumentNumber: mock.Anything,
			},
			wantErr: false,
			mock: func(repository *mocks.Account) {
				repository.On("Find", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      10,
					DocumentNumber: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "should return error when repository returns error",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				accountID:     mock.Anything,
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.Account) {
				repository.On("Find", mock.Anything, mock.Anything).Return(nil, errors.New("error to find account")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.repository)

			r := recoverAccount{
				repository: tt.fields.repository,
			}

			got, err := r.Execute(tt.args.parentContext, tt.args.accountID)

			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v")
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Execute() = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
