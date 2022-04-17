// Code generated by mockery v2.10.4. DO NOT EDIT.

package mock

import (
	model "neural_storage/cube/core/entities/model"

	mock "github.com/stretchr/testify/mock"

	user "neural_storage/cube/core/entities/user"
)

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// ValidateModelInfo provides a mock function with given fields: info
func (_m *Validator) ValidateModelInfo(info *model.Info) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Info) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateUserInfo provides a mock function with given fields: info
func (_m *Validator) ValidateUserInfo(info *user.Info) bool {
	ret := _m.Called(info)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*user.Info) bool); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
