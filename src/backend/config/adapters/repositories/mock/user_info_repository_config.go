// Code generated by mockery v2.11.0. DO NOT EDIT.

package mock

import (
	database "neural_storage/database/core/services/interactor/database"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// UserInfoRepositoryConfig is an autogenerated mock type for the UserInfoRepositoryConfig type
type UserInfoRepositoryConfig struct {
	mock.Mock
}

// Adapter provides a mock function with given fields:
func (_m *UserInfoRepositoryConfig) Adapter() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ConnParams provides a mock function with given fields:
func (_m *UserInfoRepositoryConfig) ConnParams() database.Params {
	ret := _m.Called()

	var r0 database.Params
	if rf, ok := ret.Get(0).(func() database.Params); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(database.Params)
	}

	return r0
}

// IsMocked provides a mock function with given fields:
func (_m *UserInfoRepositoryConfig) IsMocked() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewUserInfoRepositoryConfig creates a new instance of UserInfoRepositoryConfig. It also registers a cleanup function to assert the mocks expectations.
func NewUserInfoRepositoryConfig(t testing.TB) *UserInfoRepositoryConfig {
	mock := &UserInfoRepositoryConfig{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
