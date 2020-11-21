// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// CreateTransaction is an autogenerated mock type for the CreateTransaction type
type CreateTransaction struct {
	mock.Mock
}

// Execute provides a mock function with given fields: parentContext, transaction
func (_m *CreateTransaction) Execute(parentContext context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	ret := _m.Called(parentContext, transaction)

	var r0 *model.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *model.Transaction) *model.Transaction); ok {
		r0 = rf(parentContext, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.Transaction) error); ok {
		r1 = rf(parentContext, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}