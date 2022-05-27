//go:build unit
// +build unit

package modelstructweightsinfo

import (
	"neural_storage/cube/core/entities/neuron"
	"neural_storage/cube/core/entities/neuron/link"
	"neural_storage/cube/core/entities/neuron/link/weight"
	"neural_storage/cube/core/entities/neuron/offset"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/entities/structure/layer"
	"neural_storage/cube/core/entities/structure/weights"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbweight "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

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
	id := "weights 1"

	structureInfo := structure.NewInfo(
		"awesome struct",
		[]*neuron.Info{neuron.NewInfo("neuron1", "test")},
		[]*layer.Info{layer.NewInfo("test", "alpha", "beta")},
		[]*link.Info{link.NewInfo("link1", "neuron1", "neuron1")},
		[]*weights.Info{
			weights.NewInfo(
				id,
				"awesome_struct",
				[]*weight.Info{weight.NewInfo(id, "w1", 0.1)},
				[]*offset.Info{offset.NewInfo(id, "o1", 0.5)},
			),
		},
	)

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "weights_info" WHERE id = .*$`).
		WillReturnRows(utils.MockRows(dbweights.Weights{
			ID:   structureInfo.Weights()[0].ID(),
			Name: structureInfo.Weights()[0].Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "neuron_offsets" WHERE weights_id = .*$`).
		WillReturnRows(utils.MockRows(dboffset.Offset{
			Weights: structureInfo.Weights()[0].ID(),
			ID:      structureInfo.Weights()[0].Offsets()[0].ID(),
			Neuron:  structureInfo.Weights()[0].Offsets()[0].NeuronID(),
			Offset:  structureInfo.Weights()[0].Offsets()[0].Offset()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "link_weights" WHERE weights_id = .*$`).
		WillReturnRows(utils.MockRows(dbweight.Weight{
			WeightsID: structureInfo.Weights()[0].ID(),
			ID:        structureInfo.Weights()[0].Weights()[0].ID(),
			LinkID:    structureInfo.Weights()[0].Weights()[0].LinkID(),
			Value:     structureInfo.Weights()[0].Weights()[0].Weight()}))

	res, err := s.repo.Get("weights1")

	require.NoError(s.T(), err)
	require.Equal(s.T(), structureInfo.Weights()[0], res)
}

func TestGetSuite(t *testing.T) {
	suite.Run(t, new(GetSuite))
}
