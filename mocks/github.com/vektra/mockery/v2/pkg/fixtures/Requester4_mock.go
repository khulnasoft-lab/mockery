// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Requester4 is an autogenerated mock type for the Requester4 type
type Requester4 struct {
	mock.Mock
}

type Requester4_Expecter struct {
	mock *mock.Mock
}

func (_m *Requester4) EXPECT() *Requester4_Expecter {
	return &Requester4_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with no fields
func (_m *Requester4) Get() {
	_m.Called()
}

// Requester4_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Requester4_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
func (_e *Requester4_Expecter) Get() *Requester4_Get_Call {
	return &Requester4_Get_Call{Call: _e.mock.On("Get")}
}

func (_c *Requester4_Get_Call) Run(run func()) *Requester4_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Requester4_Get_Call) Return() *Requester4_Get_Call {
	_c.Call.Return()
	return _c
}

func (_c *Requester4_Get_Call) RunAndReturn(run func()) *Requester4_Get_Call {
	_c.Run(run)
	return _c
}

// NewRequester4 creates a new instance of Requester4. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequester4(t interface {
	mock.TestingT
	Cleanup(func())
}) *Requester4 {
	mock := &Requester4{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
