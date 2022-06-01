// Code generated by mockery v2.11.0. DO NOT EDIT.

package mock

import (
	model "neural_storage/cube/core/entities/model"

	mock "github.com/stretchr/testify/mock"

	modelstat "neural_storage/cube/core/entities/model/modelstat"

	repositories "neural_storage/cube/core/ports/repositories"

	structure "neural_storage/cube/core/entities/structure"

	testing "testing"

	time "time"
)

// ModelInfoRepository is an autogenerated mock type for the ModelInfoRepository type
type ModelInfoRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: info
func (_m *ModelInfoRepository) Add(info model.Info) (string, error) {
	ret := _m.Called(info)

	var r0 string
	if rf, ok := ret.Get(0).(func(model.Info) string); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Info) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: info
func (_m *ModelInfoRepository) Delete(info model.Info) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Info) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: filter
func (_m *ModelInfoRepository) Find(filter repositories.ModelInfoFilter) ([]*model.Info, error) {
	ret := _m.Called(filter)

	var r0 []*model.Info
	if rf, ok := ret.Get(0).(func(repositories.ModelInfoFilter) []*model.Info); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(repositories.ModelInfoFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: modelId
func (_m *ModelInfoRepository) Get(modelId string) (*model.Info, error) {
	ret := _m.Called(modelId)

	var r0 *model.Info
	if rf, ok := ret.Get(0).(func(string) *model.Info); ok {
		r0 = rf(modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddStat provides a mock function with given fields: from, to
func (_m *ModelInfoRepository) GetAddStat(from time.Time, to time.Time) ([]*modelstat.Info, error) {
	ret := _m.Called(from, to)

	var r0 []*modelstat.Info
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []*modelstat.Info); ok {
		r0 = rf(from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*modelstat.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStructure provides a mock function with given fields: modelId
func (_m *ModelInfoRepository) GetStructure(modelId string) (*structure.Info, error) {
	ret := _m.Called(modelId)

	var r0 *structure.Info
	if rf, ok := ret.Get(0).(func(string) *structure.Info); ok {
		r0 = rf(modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structure.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdateStat provides a mock function with given fields: from, to
func (_m *ModelInfoRepository) GetUpdateStat(from time.Time, to time.Time) ([]*modelstat.Info, error) {
	ret := _m.Called(from, to)

	var r0 []*modelstat.Info
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) []*modelstat.Info); ok {
		r0 = rf(from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*modelstat.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: info
func (_m *ModelInfoRepository) Update(info model.Info) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Info) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewModelInfoRepository creates a new instance of ModelInfoRepository. It also registers a cleanup function to assert the mocks expectations.
func NewModelInfoRepository(t testing.TB) *ModelInfoRepository {
	mock := &ModelInfoRepository{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
