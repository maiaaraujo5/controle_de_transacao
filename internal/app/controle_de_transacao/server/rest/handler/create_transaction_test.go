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
	"time"
)

type CreateTransactionSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestCreateTransactionSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionSuite))
}

func (s *CreateTransactionSuite) SetupSuite() {
	s.echo = echo.New()
}

func (s *CreateTransactionSuite) TestNewCreateTransaction() {
	ser := new(mocks.CreateTransaction)
	validate := validator.New()

	type args struct {
		service  service.CreateTransaction
		validate *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want *CreateTransaction
	}{
		{
			name: "should success build NewCreateAccount",
			args: args{
				service:  ser,
				validate: validate,
			},
			want: &CreateTransaction{
				service:  ser,
				validate: validate,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewCreateTransaction(tt.args.service, tt.args.validate)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewCreateAccount() = %v, want %v", got, tt.want)
		})
	}
}

func (s *CreateTransactionSuite) TestCreateTransaction_Handle() {

	type fields struct {
		service  *mocks.CreateTransaction
		validate *validator.Validate
	}
	type args struct {
		body io.Reader
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(service *mocks.CreateTransaction)
	}{
		{
			name: "should success create a new transaction",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "operation_type_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusCreated,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          128.50,
					EventDate:       time.Date(2020, time.November, 20, 20, 07, 35, 0, time.UTC),
				}, nil)
			},
		},
		{
			name: "should return bad request when body don't have account_id",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"operation_type_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Maybe()
			},
		},
		{
			name: "should return bad request when body don't have operation_type_id",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Maybe()
			},
		},
		{
			name: "should return bad request when body don't have amount",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "operation_type_id":1}`),
			},
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Maybe()
			},
		},
		{
			name: "should return bad request when body is malformed ",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, operation_type_id:1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Maybe()
			},
		},
		{
			name: "should return internal server error when service return unmapped error",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "operation_type_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusInternalServerError,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("error to create new transaction"))
			},
		},
		{
			name: "should return bad request when service return bad request error",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "operation_type_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusBadRequest,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors2.BadRequest("bad request"))
			},
		},
		{
			name: "should return not found when service return not found error",
			fields: fields{
				service:  new(mocks.CreateTransaction),
				validate: validator.New(),
			},
			args: args{
				body: strings.NewReader(`{"account_id":1, "operation_type_id":1, "amount":128.50}`),
			},
			wantHttpStatusCode: http.StatusNotFound,
			mock: func(service *mocks.CreateTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors2.NotFound("not found"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			tt.mock(tt.fields.service)

			h := &CreateTransaction{
				service:  tt.fields.service,
				validate: tt.fields.validate,
			}

			req := httptest.NewRequest(http.MethodPost, "/transactions", tt.args.body)
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
