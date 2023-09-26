// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import mock "github.com/stretchr/testify/mock"

// MockRelationObserver is an autogenerated mock type for the RelationObserver type
type MockRelationObserver struct {
	mock.Mock
}

// Action provides a mock function with given fields:
func (_m *MockRelationObserver) Action() {
	_m.Called()
}

// GetAttackerName provides a mock function with given fields:
func (_m *MockRelationObserver) GetAttackerName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewMockRelationObserver interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRelationObserver creates a new instance of MockRelationObserver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRelationObserver(t mockConstructorTestingTNewMockRelationObserver) *MockRelationObserver {
	mock := &MockRelationObserver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
