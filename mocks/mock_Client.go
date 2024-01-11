// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MockClient is an autogenerated mock type for the Client type
type MockClient[C interface{}] struct {
	mock.Mock
}

type MockClient_Expecter[C interface{}] struct {
	mock *mock.Mock
}

func (_m *MockClient[C]) EXPECT() *MockClient_Expecter[C] {
	return &MockClient_Expecter[C]{mock: &_m.Mock}
}

// Transaction provides a mock function with given fields: fn
func (_m *MockClient[C]) Transaction(fn func(C) error) error {
	ret := _m.Called(fn)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(C) error) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type MockClient_Transaction_Call[C interface{}] struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - fn func(C) error
func (_e *MockClient_Expecter[C]) Transaction(fn interface{}) *MockClient_Transaction_Call[C] {
	return &MockClient_Transaction_Call[C]{Call: _e.mock.On("Transaction", fn)}
}

func (_c *MockClient_Transaction_Call[C]) Run(run func(fn func(C) error)) *MockClient_Transaction_Call[C] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(C) error))
	})
	return _c
}

func (_c *MockClient_Transaction_Call[C]) Return(_a0 error) *MockClient_Transaction_Call[C] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Transaction_Call[C]) RunAndReturn(run func(func(C) error) error) *MockClient_Transaction_Call[C] {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient[C interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient[C] {
	mock := &MockClient[C]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
