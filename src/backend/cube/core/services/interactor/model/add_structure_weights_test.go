//go:build unit
// +build unit

package model

import (
	"neural_storage/cube/core/entities/structure"
	sw "neural_storage/cube/core/entities/structure/weights"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AddStructureWeightsSuite struct {
	TestSuite
}

func (s *AddStructureWeightsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *AddStructureWeightsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *AddStructureWeightsSuite) TestAdd() {
	m := sw.NewInfo("", "", nil, nil)

	s.mockedModelInfo.
		On("GetStructure", mock.Anything).
		Return(structure.NewInfo("", "", nil, nil, nil, nil), nil)

	s.mockedValidator.On("ValidateModelInfo", mock.Anything).Return(nil)

	s.mockedWeightsInfo.On("Add", mock.Anything, mock.Anything).Return(nil, nil)
	_, err := s.interactor.AddStructureWeights(s.ctx, "", "", *m)

	require.NoError(s.T(), err)
}

func TestAddStructureWeightsSuite(t *testing.T) {
	suite.Run(t, new(AddStructureWeightsSuite))
}
