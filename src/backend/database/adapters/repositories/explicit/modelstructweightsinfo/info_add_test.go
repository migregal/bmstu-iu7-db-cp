//go:build unit
// +build unit

package modelstructweightsinfo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure/weights"
	dbneuron "neural_storage/database/core/entities/neuron"
	dblink "neural_storage/database/core/entities/neuron/link"
	dbweights "neural_storage/database/core/entities/structure/weights"
	dblayer "neural_storage/database/core/entities/structure/layer"
	"neural_storage/database/test/mock/utils"
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
		[]*weight.Info{weight.NewInfo(1, 1, 10)},
		[]*offset.Info{offset.NewInfo(1, 1, 0.1)},
	)

	s.SqlMock.ExpectBegin()
	s.SqlMock.ExpectQuery(`^SELECT \* FROM "layers" WHERE structure_id = .*$`).WillReturnRows(utils.MockRows(dblayer.Layer{ID: 1}))
	s.SqlMock.ExpectQuery(`^SELECT \* FROM "neurons" WHERE layer_id in .*$`).WillReturnRows(utils.MockRows(dbneuron.Neuron{ID: 0, LayerID: 0}))
	s.SqlMock.ExpectQuery(`^SELECT \* FROM "neuron_links" WHERE from_id in .*$`).WillReturnRows(utils.MockRows(dblink.Link{}))
	s.SqlMock.ExpectQuery(`^INSERT INTO "weights_info" .* RETURNING "id"$`).WillReturnRows(utils.MockRows(dbweights.Weights{ID: 1}))
	s.SqlMock.ExpectExec(`^INSERT INTO "neuron_offsets" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectExec(`^INSERT INTO "link_weights" .*$`).WillReturnResult(sqlmock.NewResult(1, 0))
	s.SqlMock.ExpectCommit()

	_, err := s.repo.Add("awesome_struct_id", []weights.Info{*info})

	require.NoError(s.T(), err)
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
