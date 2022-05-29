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

type DeleteStructureWeightsSuite struct {
	TestSuite
}

func (s *DeleteStructureWeightsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *DeleteStructureWeightsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *DeleteStructureWeightsSuite) TestDelete() {
	m := sw.NewInfo("", "", nil, nil)
	s.mockedWeightsInfo.On("Delete", mock.Anything, mock.Anything).Return(nil)

	err := s.interactor.DeleteStructureWeights(s.ctx, "test", m.ID())

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedWeightsInfo.AssertExpectations(s.T()))
}

func TestDeleteStructureWeightsSuite(t *testing.T) {
	suite.Run(t, new(DeleteStructureWeightsSuite))
}
