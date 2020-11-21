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
	"time"
)

type RecoverTransactionSuite struct {
	suite.Suite
	echo *echo.Echo
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(RecoverTransactionSuite))
}

func (s *RecoverTransactionSuite) SetupSuite() {
	s.echo = echo.New()
}

func (s *RecoverTransactionSuite) TestNewRecoverTransaction() {

	serv := new(mocks.RecoverTransaction)
	type args struct {
		service service.RecoverTransaction
	}
	tests := []struct {
		name string
		args args
		want *RecoverTransaction
	}{
		{
			name: "should success build NewRecoverTransaction",
			args: args{
				service: serv,
			},
			want: &RecoverTransaction{
				service: serv,
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			got := NewRecoverTransaction(tt.args.service)
			s.Assert().True(reflect.DeepEqual(got, tt.want), "NewRecoverTransaction() = %v, want %v", got, tt.want)
		})
	}
}

func (s *RecoverTransactionSuite) TestRecoverTransaction_Handle() {
	type fields struct {
		service *mocks.RecoverTransaction
	}
	type args struct {
		transactionID string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(service *mocks.RecoverTransaction)
	}{
		{
			name: "should success recover one transaction",
			fields: fields{
				service: new(mocks.RecoverTransaction),
			},
			args: args{
				transactionID: mock.Anything,
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(service *mocks.RecoverTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(&model.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -100,
					EventDate:       time.Date(2020, time.November, 20, 20, 07, 35, 0, time.UTC),
				}, nil)
			},
		},
		{
			name: "should return error when service return error",
			fields: fields{
				service: new(mocks.RecoverTransaction),
			},
			args: args{
				transactionID: mock.Anything,
			},
			wantHttpStatusCode: http.StatusInternalServerError,
			mock: func(service *mocks.RecoverTransaction) {
				service.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("error to recover transaction"))
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			tt.mock(tt.fields.service)

			h := &RecoverTransaction{
				service: tt.fields.service,
			}

			req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
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
