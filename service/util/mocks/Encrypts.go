// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Encrypts is an autogenerated mock type for the Encrypts type
type Encrypts struct {
	mock.Mock
}

// DecryptAES provides a mock function with given fields: stringToDecrypt, passPhrase
func (_m *Encrypts) DecryptAES(stringToDecrypt string, passPhrase string) (string, error) {
	ret := _m.Called(stringToDecrypt, passPhrase)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(stringToDecrypt, passPhrase)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(stringToDecrypt, passPhrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EncryptAES provides a mock function with given fields: stringToEncrypt, passPhrase
func (_m *Encrypts) EncryptAES(stringToEncrypt string, passPhrase string) (string, error) {
	ret := _m.Called(stringToEncrypt, passPhrase)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(stringToEncrypt, passPhrase)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(stringToEncrypt, passPhrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEncrypts interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncrypts creates a new instance of Encrypts. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncrypts(t mockConstructorTestingTNewEncrypts) *Encrypts {
	mock := &Encrypts{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
