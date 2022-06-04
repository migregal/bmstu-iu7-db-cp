//go:build unit
// +build unit

package model

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
	sw "neural_storage/cube/core/entities/structure/weights"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UpdateStructureWeightsSuite struct {
	TestSuite
}

func (s *UpdateStructureWeightsSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *UpdateStructureWeightsSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *UpdateStructureWeightsSuite) TestUpdate() {
	s.mockedModelInfo.
		On("Get", mock.Anything).
		Return(model.NewInfo("", "", "", structure.NewInfo("", "", nil, nil, nil, nil)), nil)

	s.mockedValidator.On("ValidateModelInfo", mock.Anything).Return(nil)

	s.mockedWeightsInfo.On("Update", mock.Anything, mock.Anything).Return(nil)

	expected := sw.NewInfo("", "", nil, nil)

	err := s.interactor.UpdateStructureWeights(s.ctx, "", "", *expected)

	require.NoError(s.T(), err)

	require.True(s.T(), s.mockedWeightsInfo.AssertExpectations(s.T()))
	require.True(s.T(), s.mockedModelInfo.AssertExpectations(s.T()))
}

func TestUpdateStructureWeightsSuite(t *testing.T) {
	suite.Run(t, new(UpdateStructureWeightsSuite))
}
