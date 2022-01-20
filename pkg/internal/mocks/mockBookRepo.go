// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	storage "github.com/vivek080/library-app/pkg/storage"
)

// MockBookRepo is an autogenerated mock type for the Repository type
type MockBookRepo struct {
	mock.Mock
}

// GetBookList provides a mock function with given fields:
func (_m *MockBookRepo) GetBookList() ([]storage.BookDetails, error) {
	ret := _m.Called()

	var r0 []storage.BookDetails
	if rf, ok := ret.Get(0).(func() []storage.BookDetails); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.BookDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveBook provides a mock function with given fields: _a0
func (_m *MockBookRepo) SaveBook(_a0 *storage.BookDetails) (*storage.BookDetails, error) {
	ret := _m.Called(_a0)

	var r0 *storage.BookDetails
	if rf, ok := ret.Get(0).(func(*storage.BookDetails) *storage.BookDetails); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.BookDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*storage.BookDetails) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
