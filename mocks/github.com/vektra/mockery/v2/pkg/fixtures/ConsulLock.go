// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConsulLock is an autogenerated mock type for the ConsulLock type
type ConsulLock struct {
	mock.Mock
}

type ConsulLock_Expecter struct {
	mock *mock.Mock
}

func (_m *ConsulLock) EXPECT() *ConsulLock_Expecter {
	return &ConsulLock_Expecter{mock: &_m.Mock}
}

// Lock provides a mock function with given fields: _a0
func (_m *ConsulLock) Lock(_a0 <-chan struct{}) (<-chan struct{}, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("Missing Return() function for Lock")
	}

	var r0 <-chan struct{}
	var r1 error
	if rf, ok := ret.Get(0).(func(<-chan struct{}) (<-chan struct{}, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(<-chan struct{}) <-chan struct{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(<-chan struct{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConsulLock_Lock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Lock'
type ConsulLock_Lock_Call struct {
	*mock.Call
}

// Lock is a helper method to define mock.On call
//   - _a0 <-chan struct{}
func (_e *ConsulLock_Expecter) Lock(_a0 interface{}) *ConsulLock_Lock_Call {
	return &ConsulLock_Lock_Call{Call: _e.mock.On("Lock", _a0)}
}

func (_c *ConsulLock_Lock_Call) Run(run func(_a0 <-chan struct{})) *ConsulLock_Lock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(<-chan struct{}))
	})
	return _c
}

func (_c *ConsulLock_Lock_Call) Return(_a0 <-chan struct{}, _a1 error) *ConsulLock_Lock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConsulLock_Lock_Call) RunAndReturn(run func(<-chan struct{}) (<-chan struct{}, error)) *ConsulLock_Lock_Call {
	_c.Call.Return(run)
	return _c
}

// Unlock provides a mock function with given fields:
func (_m *ConsulLock) Unlock() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("Missing Return() function for Unlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsulLock_Unlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unlock'
type ConsulLock_Unlock_Call struct {
	*mock.Call
}

// Unlock is a helper method to define mock.On call
func (_e *ConsulLock_Expecter) Unlock() *ConsulLock_Unlock_Call {
	return &ConsulLock_Unlock_Call{Call: _e.mock.On("Unlock")}
}

func (_c *ConsulLock_Unlock_Call) Run(run func()) *ConsulLock_Unlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConsulLock_Unlock_Call) Return(_a0 error) *ConsulLock_Unlock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConsulLock_Unlock_Call) RunAndReturn(run func() error) *ConsulLock_Unlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsulLock creates a new instance of ConsulLock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsulLock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsulLock {
	mock := &ConsulLock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
