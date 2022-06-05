//go:build unit
// +build unit

package modelinfo

import (
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DeleteSuite struct {
	TestSuite
}

func (s *DeleteSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *DeleteSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *DeleteSuite) TestDelete() {
	name := "test"
	info := model.NewInfo(
		"id",
		"",
		name,
		structure.NewInfo(
			"",
			"awesome struct",
			[]*neuron.Info{neuron.NewInfo(1, 1)},
			[]*layer.Info{layer.NewInfo(1, "alpha", "beta")},
			[]*link.Info{link.NewInfo(1, 1, 1)},
			[]*weights.Info{
				weights.NewInfo(
					"",
					"weights1",
					[]*weight.Info{weight.NewInfo(1, 1, 0.1)},
					[]*offset.Info{offset.NewInfo(1, 1, 0.5)},
				),
			},
		))

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectExec(`^DELETE FROM "models" WHERE "models"."id" = .*$`).WillReturnResult(sqlmock.NewResult(0, 1))
	s.SqlMock.ExpectCommit()

	err := s.repo.Delete(*info)

	require.NoError(s.T(), err)
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
