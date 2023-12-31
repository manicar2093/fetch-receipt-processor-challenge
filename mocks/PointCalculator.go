// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	receipts "github.com/manicar2093/fetch-receipt-processor-challenge/internal/receipts"
	mock "github.com/stretchr/testify/mock"
)

// PointCalculator is an autogenerated mock type for the PointCalculator type
type PointCalculator struct {
	mock.Mock
}

type PointCalculator_Expecter struct {
	mock *mock.Mock
}

func (_m *PointCalculator) EXPECT() *PointCalculator_Expecter {
	return &PointCalculator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *PointCalculator) Execute(_a0 receipts.Receipt) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(receipts.Receipt) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// PointCalculator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type PointCalculator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 receipts.Receipt
func (_e *PointCalculator_Expecter) Execute(_a0 interface{}) *PointCalculator_Execute_Call {
	return &PointCalculator_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *PointCalculator_Execute_Call) Run(run func(_a0 receipts.Receipt)) *PointCalculator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(receipts.Receipt))
	})
	return _c
}

func (_c *PointCalculator_Execute_Call) Return(_a0 int) *PointCalculator_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PointCalculator_Execute_Call) RunAndReturn(run func(receipts.Receipt) int) *PointCalculator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPointCalculator interface {
	mock.TestingT
	Cleanup(func())
}

// NewPointCalculator creates a new instance of PointCalculator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPointCalculator(t mockConstructorTestingTNewPointCalculator) *PointCalculator {
	mock := &PointCalculator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
