// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UUIDGen is an autogenerated mock type for the UUIDGen type
type UUIDGen struct {
	mock.Mock
}

type UUIDGen_Expecter struct {
	mock *mock.Mock
}

func (_m *UUIDGen) EXPECT() *UUIDGen_Expecter {
	return &UUIDGen_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *UUIDGen) Execute() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// UUIDGen_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type UUIDGen_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *UUIDGen_Expecter) Execute() *UUIDGen_Execute_Call {
	return &UUIDGen_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *UUIDGen_Execute_Call) Run(run func()) *UUIDGen_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UUIDGen_Execute_Call) Return(_a0 uuid.UUID) *UUIDGen_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UUIDGen_Execute_Call) RunAndReturn(run func() uuid.UUID) *UUIDGen_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUUIDGen interface {
	mock.TestingT
	Cleanup(func())
}

// NewUUIDGen creates a new instance of UUIDGen. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUUIDGen(t mockConstructorTestingTNewUUIDGen) *UUIDGen {
	mock := &UUIDGen{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}