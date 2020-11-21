package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type RecoverAccountSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestRecoverAccount(t *testing.T) {
	suite.Run(t, new(RecoverAccountSuite))
}

func (s *RecoverAccountSuite) SetupSuite() {
	s.echo = echo.New()
}

func (s *RecoverAccountSuite) TestNewRecoverAccount() {

	ser := new(mocks.RecoverAccount)

	type args struct {
		service service.RecoverAccount
	}
	tests := []struct {
		name string
		args args
		want *RecoverAccount
	}{
		{
			name: "should success build NewRecoverAccount",
			args: args{
				service: ser,
			},
			want: &RecoverAccount{
				service: ser,
			},
		},
	}
	for _, tt := range tests {

		s.Run(tt.name, func() {
			got := NewRecoverAccount(tt.args.service)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRecoverAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *RecoverAccountSuite) TestRecoverAccount_Handle() {
	type fields struct {
		service *mocks.RecoverAccount
	}
	type args struct {
		accountID string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(service *mocks.RecoverAccount)
	}{
		{
			name: "should success recover one account",
			fields: fields{
				service: new(mocks.RecoverAccount),
			},
			args: args{
				accountID: mock.Anything,
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(service *mocks.RecoverAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      1,
					DocumentNumber: mock.Anything,
				}, nil).Once()
			},
		},
		{
			name: "should return internal server error when service return error",
			fields: fields{
				service: new(mocks.RecoverAccount),
			},
			args: args{
				accountID: mock.Anything,
			},
			wantHttpStatusCode: http.StatusInternalServerError,
			mock: func(service *mocks.RecoverAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("error to recover account")).Once()
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.service)

			h := &RecoverAccount{
				service: tt.fields.service,
			}

			req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
			rec := httptest.NewRecorder()

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			context := s.echo.NewContext(req, rec)

			err := h.Handle(context)

			if err != nil {
				log.Fatal(err)
			}

			s.Assert().Equal(tt.wantHttpStatusCode, context.Response().Status)
			tt.fields.service.AssertExpectations(s.T())
		})

	}
}
