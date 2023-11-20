// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RequesterSlice is an autogenerated mock type for the RequesterSlice type
type RequesterSlice struct {
	mock.Mock
}

type RequesterSlice_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterSlice) EXPECT() *RequesterSlice_Expecter {
	return &RequesterSlice_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: path
func (_m *RequesterSlice) Get(path string) ([]string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("Missing Return() function for Get")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequesterSlice_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterSlice_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path string
func (_e *RequesterSlice_Expecter) Get(path interface{}) *RequesterSlice_Get_Call {
	return &RequesterSlice_Get_Call{Call: _e.mock.On("Get", path)}
}

func (_c *RequesterSlice_Get_Call) Run(run func(path string)) *RequesterSlice_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequesterSlice_Get_Call) Return(_a0 []string, _a1 error) *RequesterSlice_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequesterSlice_Get_Call) RunAndReturn(run func(string) ([]string, error)) *RequesterSlice_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequesterSlice creates a new instance of RequesterSlice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterSlice(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterSlice {
	mock := &RequesterSlice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
