// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"

// PageableStore is an autogenerated mock type for the PageableStore type
type PageableStore struct {
	mock.Mock
}

// GetByKey provides a mock function with given fields: key
func (_m *PageableStore) GetByKey(key string) (interface{}, bool, error) {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// List provides a mock function with given fields:
func (_m *PageableStore) List() []interface{} {
	ret := _m.Called()

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func() []interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// ListKeys provides a mock function with given fields:
func (_m *PageableStore) ListKeys() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}
