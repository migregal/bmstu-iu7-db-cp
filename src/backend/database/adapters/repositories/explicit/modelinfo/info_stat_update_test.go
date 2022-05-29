//go:build unit
// +build unit

package modelinfo

import (
	"neural_storage/cube/core/entities/model/modelstat"
	dbstat "neural_storage/database/core/entities/model/modelstat"
	"neural_storage/database/test/mock/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetUpdateStatSuite struct {
	TestSuite
}

func (s *GetUpdateStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetUpdateStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetUpdateStatSuite) TestGet() {

	info := []*modelstat.Info{modelstat.New("test", time.Now(), time.Now())}

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "models" WHERE updated_at > .* AND updated_at < .*$`).
		WillReturnRows(utils.MockRows(
			dbstat.Info{
				ID:        info[0].ID(),
				CreatedAt: info[0].CreatedAt(),
				UpdatedAt: info[0].UpdatedAt(),
			}),
		)

	res, err := s.repo.GetUpdateStat(time.Now(), time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, res)
}

func TestGetUpdateStatSuite(t *testing.T) {
	suite.Run(t, new(GetUpdateStatSuite))
}
