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
	"neural_storage/cube/core/ports/repositories"
	dboffset "neural_storage/database/core/entities/neuron/offset"
	dbweight "neural_storage/database/core/entities/structure/weight"
	dbweights "neural_storage/database/core/entities/structure/weights"
	"neural_storage/database/test/mock/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FindSuite struct {
	TestSuite
}

func (s *FindSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *FindSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *FindSuite) TestFind() {
	name := "test"
	structureInfo := structure.NewInfo(
		"",
		"awesome struct",
		[]*neuron.Info{neuron.NewInfo(1, 1)},
		[]*layer.Info{layer.NewInfo(1, "alpha", "beta")},
		[]*link.Info{link.NewInfo(1, 1, 1)},
		[]*weights.Info{
			weights.NewInfo(
				"",
				name,
				[]*weight.Info{weight.NewInfo(1, 1, 0.1)},
				[]*offset.Info{offset.NewInfo(1, 1, 0.5)},
			),
		},
	)
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "weights_info" WHERE id in .*$`).
		WillReturnRows(utils.MockRows(dbweights.Weights{
			InnerID:   structureInfo.Weights()[0].ID(),
			Name: structureInfo.Weights()[0].Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "weights_info" WHERE id = .* ORDER BY .* LIMIT 1$`).
		WillReturnRows(utils.MockRows(dbweights.Weights{
			InnerID:   structureInfo.Weights()[0].ID(),
			Name: structureInfo.Weights()[0].Name()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "neuron_offsets" WHERE weights_info_id = .*$`).
		WillReturnRows(utils.MockRows(dboffset.Offset{
			InnerWeights: structureInfo.Weights()[0].ID(),
			ID:      structureInfo.Weights()[0].Offsets()[0].ID(),
			Neuron:  structureInfo.Weights()[0].Offsets()[0].NeuronID(),
			Offset:  structureInfo.Weights()[0].Offsets()[0].Offset()}))
	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "link_weights" WHERE weights_info_id = .*$`).
		WillReturnRows(utils.MockRows(dbweight.Weight{
			InnerWeightsID: structureInfo.Weights()[0].ID(),
			ID:        structureInfo.Weights()[0].Weights()[0].ID(),
			LinkID:    structureInfo.Weights()[0].Weights()[0].LinkID(),
			Value:     structureInfo.Weights()[0].Weights()[0].Weight()}))

	res, err := s.repo.Find(repositories.StructWeightsInfoFilter{Ids: []string{name}, Limit: 10})

	require.NoError(s.T(), err)
	require.Equal(s.T(), structureInfo.Weights(), res)
}

func TestFindSuite(t *testing.T) {
	suite.Run(t, new(FindSuite))
}
