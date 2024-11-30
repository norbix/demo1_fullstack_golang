// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// AccountHandler is an autogenerated mock type for the AccountHandler type
type AccountHandler struct {
	mock.Mock
}

type AccountHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountHandler) EXPECT() *AccountHandler_Expecter {
	return &AccountHandler_Expecter{mock: &_m.Mock}
}

// CreateAccount provides a mock function with given fields: _a0, _a1
func (_m *AccountHandler) CreateAccount(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// AccountHandler_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type AccountHandler_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *AccountHandler_Expecter) CreateAccount(_a0 interface{}, _a1 interface{}) *AccountHandler_CreateAccount_Call {
	return &AccountHandler_CreateAccount_Call{Call: _e.mock.On("CreateAccount", _a0, _a1)}
}

func (_c *AccountHandler_CreateAccount_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *AccountHandler_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *AccountHandler_CreateAccount_Call) Return() *AccountHandler_CreateAccount_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandler_CreateAccount_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *AccountHandler_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccounts provides a mock function with given fields: _a0, _a1
func (_m *AccountHandler) GetAccounts(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// AccountHandler_GetAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccounts'
type AccountHandler_GetAccounts_Call struct {
	*mock.Call
}

// GetAccounts is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *AccountHandler_Expecter) GetAccounts(_a0 interface{}, _a1 interface{}) *AccountHandler_GetAccounts_Call {
	return &AccountHandler_GetAccounts_Call{Call: _e.mock.On("GetAccounts", _a0, _a1)}
}

func (_c *AccountHandler_GetAccounts_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *AccountHandler_GetAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *AccountHandler_GetAccounts_Call) Return() *AccountHandler_GetAccounts_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandler_GetAccounts_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *AccountHandler_GetAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// NewAccountHandler creates a new instance of AccountHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountHandler {
	mock := &AccountHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
