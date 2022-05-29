//go:build unit
// +build unit

package model

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sw "neural_storage/cube/core/entities/structure/weights"
	"neural_storage/cube/core/ports/interactors"
)

type FindStructureWeightsSuite struct {
	TestSuite
}

func (s *FindStructureWeightsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *FindStructureWeightsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *FindStructureWeightsSuite) TestFind() {
	filter := interactors.ModelWeightsInfoFilter{}
	expected := []*sw.Info{
		sw.NewInfo("", "", nil, nil),
	}

	s.mockedWeightsInfo.On("Find", mock.Anything).Return(expected, nil)
	info, err := s.interactor.FindStructureWeights(s.ctx, filter)

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, expected)

	require.True(s.T(), s.mockedWeightsInfo.AssertExpectations(s.T()))
}

func TestFindStructureWeightsSuite(t *testing.T) {
	suite.Run(t, new(FindStructureWeightsSuite))
}
