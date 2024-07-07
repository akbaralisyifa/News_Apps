// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	users "newsapps/internal/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// Login provides a mock function with given fields: email
func (_m *Query) Login(email string) (users.Users, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 users.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (users.Users, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) users.Users); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(users.Users)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *Query) Register(newUser users.Users) error {
	ret := _m.Called(newUser)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(users.Users) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuery creates a new instance of Query. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *Query {
	mock := &Query{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
