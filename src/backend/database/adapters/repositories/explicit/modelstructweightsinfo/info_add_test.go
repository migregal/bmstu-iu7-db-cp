//go:build unit
// +build unit

package modelstructweightsinfo

import (
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
	info := weights.NewInfo(
		"awesome_id",
		"test",
		[]*weight.Info{weight.NewInfo("weight 1", "link 1", 10)},
		[]*offset.Info{offset.NewInfo("offset 1", "neuron 1", 0.1)},
	)

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectQuery(`^INSERT INTO "weights_info" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbweights.Weights{ID: "some id for weights"}))
	s.SqlMock.ExpectExec(`^INSERT INTO "neuron_offsets" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectExec(`^INSERT INTO "link_weights" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectCommit()

	err := s.repo.Add("awesome_struct_id", []weights.Info{*info})

	require.NoError(s.T(), err)
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
