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

type GetSuite struct {
	TestSuite
}

func (s *GetSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetSuite) TestGet() {
	expected := model.NewInfo("", "", "", nil)

	s.mockedModelInfo.On("Get", mock.Anything).Return(expected, nil)
	info, err := s.interactor.Get(s.ctx, expected.ID())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
