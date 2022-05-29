//go:build unit
// +build unit

package model

import (
	sw "neural_storage/cube/core/entities/structure/weights"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetStructureWeightsSuite struct {
	TestSuite
}

func (s *GetStructureWeightsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetStructureWeightsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetStructureWeightsSuite) TestGet() {
	expected := sw.NewInfo("", "", nil, nil)
	s.mockedWeightsInfo.On("Get", mock.Anything).Return(expected, nil)

	info, err := s.interactor.GetStructureWeights(s.ctx, "")

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedWeightsInfo.AssertExpectations(s.T()))
}

func TestGetStructureWeightsSuite(t *testing.T) {
	suite.Run(t, new(GetStructureWeightsSuite))
}
