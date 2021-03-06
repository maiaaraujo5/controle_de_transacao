// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/domain/model"
	mock "github.com/stretchr/testify/mock"
)

// RecoverAccount is an autogenerated mock type for the RecoverAccount type
type RecoverAccount struct {
	mock.Mock
}

// Execute provides a mock function with given fields: parentContext, accountID
func (_m *RecoverAccount) Execute(parentContext context.Context, accountID string) (*model.Account, error) {
	ret := _m.Called(parentContext, accountID)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Account); ok {
		r0 = rf(parentContext, accountID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(parentContext, accountID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
