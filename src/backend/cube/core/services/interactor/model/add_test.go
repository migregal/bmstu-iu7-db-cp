//go:build unit
// +build unit

package model

import (
	"neural_storage/cube/core/entities/model"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AddSuite struct {
	TestSuite
}

func (s *AddSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *AddSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *AddSuite) TestAdd() {
	s.mockedValidator.On("ValidateModelInfo", mock.Anything).Return(nil)

	m := model.NewInfo("", "", "", nil)
	s.mockedModelInfo.On("Add", mock.Anything).Return("", nil)

	_, err := s.interactor.Add(s.ctx, *m)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
