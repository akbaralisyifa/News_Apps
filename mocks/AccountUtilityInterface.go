// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AccountUtilityInterface is an autogenerated mock type for the AccountUtilityInterface type
type AccountUtilityInterface struct {
	mock.Mock
}

// EmailPasswordValidator provides a mock function with given fields: inputEmail, inputPw
func (_m *AccountUtilityInterface) EmailPasswordValidator(inputEmail string, inputPw string) error {
	ret := _m.Called(inputEmail, inputPw)

	if len(ret) == 0 {
		panic("no return value specified for EmailPasswordValidator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(inputEmail, inputPw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterValidator provides a mock function with given fields: inputName, inputEmail, inputPw
func (_m *AccountUtilityInterface) RegisterValidator(inputName string, inputEmail string, inputPw string) error {
	ret := _m.Called(inputName, inputEmail, inputPw)

	if len(ret) == 0 {
		panic("no return value specified for RegisterValidator")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(inputName, inputEmail, inputPw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAccountUtilityInterface creates a new instance of AccountUtilityInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountUtilityInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountUtilityInterface {
	mock := &AccountUtilityInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}