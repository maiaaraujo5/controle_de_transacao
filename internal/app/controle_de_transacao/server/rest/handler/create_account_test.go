package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/service/mocks"
	errors2 "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type CreateAccountSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestCreateAccountSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountSuite))
}

func (s *CreateAccountSuite) SetupSuite() {
	s.echo = echo.New()
}

func (s *CreateAccountSuite) TestNewCreateAccount() {
	ser := new(mocks.CreateAccount)
	validate := validator.New()

	type args struct {
		service  service.CreateAccount
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want *CreateAccount
	}{
		{
			name: "should success build NewCreateAccount",
			args: args{
				service:  ser,
				validate: validate,
			},
			want: &CreateAccount{
				service:  ser,
				validate: validate,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewCreateAccount(tt.args.service, tt.args.validate)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateAccount() = %v, want %v", got, tt.want)
		})

	}
}

func (s *CreateAccountSuite) TestCreateAccount_Handle() {
	type fields struct {
		service  *mocks.CreateAccount
		validate *validator.Validate
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantHttpStatusCode int
		mock               func(service *mocks.CreateAccount)
	}{
		{
			name: "should success create a new account",
			fields: fields{
				service:  new(mocks.CreateAccount),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number": "12345689"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusCreated,
			mock: func(service *mocks.CreateAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(&model.Account{
					AccountID:      1,
					DocumentNumber: "12345689",
				}, nil)
			},
		},
		{
			name: "should return bad request when body is malformed",
			fields: fields{
				service:  new(mocks.CreateAccount),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{document_number: "123"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, nil).Maybe()
			},
		},
		{
			name: "should return bad request when body don't have document_number",
			fields: fields{
				service:  new(mocks.CreateAccount),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, nil).Maybe()
			},
		},
		{
			name: "should return internal server error when service one unmapped error",
			fields: fields{
				service:  new(mocks.CreateAccount),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number": "12345689"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusInternalServerError,
			mock: func(service *mocks.CreateAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("error to recover account")).Once()
			},
		},
		{
			name: "should return conflict when service return already exists error",
			fields: fields{
				service:  new(mocks.CreateAccount),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"document_number": "12345689"}`),
			},
			wantErr:            false,
			wantHttpStatusCode: http.StatusConflict,
			mock: func(service *mocks.CreateAccount) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors2.AlreadyExists("already exists"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.service)

			h := &CreateAccount{
				service:  tt.fields.service,
				validate: tt.fields.validate,
			}

			req := httptest.NewRequest(http.MethodPost, "/accounts", tt.args.body)
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
