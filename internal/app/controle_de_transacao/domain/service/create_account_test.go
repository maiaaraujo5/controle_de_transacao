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

type CreateAccountSuite struct {
	suite.Suite
}

func TestNewCreateAccountSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountSuite))
}

func (s *CreateAccountSuite) TestNewCreateAccount() {

	repo := new(mocks.Account)

	type args struct {
		repository repository.Account
	}
	tests := []struct {
		name string
		args args
		want CreateAccount
	}{
		{
			name: "should success build NewCreateAccount",
			args: args{
				repository: repo,
			},
			want: &createAccount{
				repository: repo,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewCreateAccount(tt.args.repository)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateAccountSuite) Test_createAccount_Execute() {
	type fields struct {
		repository *mocks.Account
	}
	type args struct {
		parentContext context.Context
		account       *model.Account
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
			name: "should success save a new account",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want: &model.Account{
				AccountID:      1,
				DocumentNumber: mock.Anything,
			},
			wantErr: false,
			mock: func(repository *mocks.Account) {
				repository.On("Save", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      1,
					DocumentNumber: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "should return error when repository return error",
			fields: fields{
				repository: new(mocks.Account),
			},
			args: args{
				parentContext: context.Background(),
				account: &model.Account{
					DocumentNumber: mock.Anything,
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.Account) {
				repository.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save new account")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.repository)

			c := createAccount{
				repository: tt.fields.repository,
			}
			got, err := c.Execute(tt.args.parentContext, tt.args.account)
			s.Assert().True((err != nil) == tt.wantErr, "Execute() error = %v, wantErr %v")
			s.Assert().True(reflect.DeepEqual(got, tt.want), "Execute() = %v, want %v", got, tt.want)

			tt.fields.repository.AssertExpectations(s.T())
		})
	}
}
