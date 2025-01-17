// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilterString is an autogenerated mock type for the FilterString type
type FilterString struct {
	mock.Mock
}

// MakeDeletedAtIsNull provides a mock function with given fields:
func (_m *FilterString) MakeDeletedAtIsNull() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeFilterEventAtBetweenString provides a mock function with given fields: start, end
func (_m *FilterString) MakeFilterEventAtBetweenString(start int64, end int64) string {
	ret := _m.Called(start, end)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64, int64) string); ok {
		r0 = rf(start, end)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeGateGroupNameFiltersNotEqualString provides a mock function with given fields: gateGroupName
func (_m *FilterString) MakeGateGroupNameFiltersNotEqualString(gateGroupName string) string {
	ret := _m.Called(gateGroupName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(gateGroupName)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeID provides a mock function with given fields: id
func (_m *FilterString) MakeID(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeIDFilters provides a mock function with given fields: id
func (_m *FilterString) MakeIDFilters(id string) []string {
	ret := _m.Called(id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MakeNotEqualID provides a mock function with given fields: id
func (_m *FilterString) MakeNotEqualID(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewFilterString interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilterString creates a new instance of FilterString. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilterString(t mockConstructorTestingTNewFilterString) *FilterString {
	mock := &FilterString{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
