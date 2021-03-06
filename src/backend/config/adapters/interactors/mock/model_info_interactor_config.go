// Code generated by mockery v2.11.0. DO NOT EDIT.

package mock

import (
	config "neural_storage/database/core/ports/config"

	mock "github.com/stretchr/testify/mock"

	portsconfig "neural_storage/cube/core/ports/config"

	testing "testing"
)

// ModelInfoInteractorConfig is an autogenerated mock type for the ModelInfoInteractorConfig type
type ModelInfoInteractorConfig struct {
	mock.Mock
}

// ModelInfoRepoConfig provides a mock function with given fields:
func (_m *ModelInfoInteractorConfig) ModelInfoRepoConfig() config.ModelInfoRepositoryConfig {
	ret := _m.Called()

	var r0 config.ModelInfoRepositoryConfig
	if rf, ok := ret.Get(0).(func() config.ModelInfoRepositoryConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ModelInfoRepositoryConfig)
		}
	}

	return r0
}

// ModelStructureWeightInfoRepoConfig provides a mock function with given fields:
func (_m *ModelInfoInteractorConfig) ModelStructureWeightInfoRepoConfig() config.ModelStructureWeightsInfoRepositoryConfig {
	ret := _m.Called()

	var r0 config.ModelStructureWeightsInfoRepositoryConfig
	if rf, ok := ret.Get(0).(func() config.ModelStructureWeightsInfoRepositoryConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ModelStructureWeightsInfoRepositoryConfig)
		}
	}

	return r0
}

// ValidatorConfig provides a mock function with given fields:
func (_m *ModelInfoInteractorConfig) ValidatorConfig() portsconfig.ValidatorConfig {
	ret := _m.Called()

	var r0 portsconfig.ValidatorConfig
	if rf, ok := ret.Get(0).(func() portsconfig.ValidatorConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(portsconfig.ValidatorConfig)
		}
	}

	return r0
}

// NewModelInfoInteractorConfig creates a new instance of ModelInfoInteractorConfig. It also registers a cleanup function to assert the mocks expectations.
func NewModelInfoInteractorConfig(t testing.TB) *ModelInfoInteractorConfig {
	mock := &ModelInfoInteractorConfig{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
