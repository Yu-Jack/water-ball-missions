// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import mock "github.com/stretchr/testify/mock"

// MockActionStrategy is an autogenerated mock type for the ActionStrategy type
type MockActionStrategy struct {
	mock.Mock
}

// S1 provides a mock function with given fields: skillsIDs
func (_m *MockActionStrategy) S1(skillsIDs []int) int {
	ret := _m.Called(skillsIDs)

	var r0 int
	if rf, ok := ret.Get(0).(func([]int) int); ok {
		r0 = rf(skillsIDs)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// S2 provides a mock function with given fields: availableIDs, limit, list
func (_m *MockActionStrategy) S2(availableIDs []int, limit int, list string) []int {
	ret := _m.Called(availableIDs, limit, list)

	var r0 []int
	if rf, ok := ret.Get(0).(func([]int, int, string) []int); ok {
		r0 = rf(availableIDs, limit, list)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockActionStrategy interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockActionStrategy creates a new instance of MockActionStrategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockActionStrategy(t mockConstructorTestingTNewMockActionStrategy) *MockActionStrategy {
	mock := &MockActionStrategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
