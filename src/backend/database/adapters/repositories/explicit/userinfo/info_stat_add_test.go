//go:build unit
// +build unit

package userinfo

import (
	"neural_storage/cube/core/entities/user/userstat"
	"neural_storage/database/core/entities/user_info/userinfostat"
	"neural_storage/database/test/mock/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type GetAddStatSuite struct {
	TestSuite
}

func (s *GetAddStatSuite) SetupTest() {
	s.TestSuite.SetupTest()
}

func (s *GetAddStatSuite) TearDownTest() {
	s.TestSuite.TearDownTest()
}

func (s *GetAddStatSuite) TestGet() {

	info := []*userstat.Info{userstat.New("test", time.Now(), time.Now())}

	s.SqlMock.
		ExpectQuery(`^SELECT \* FROM "users_info" WHERE created_at > .* AND created_at < .*$`).
		WillReturnRows(utils.MockRows(
			userinfostat.Info{
				ID:        info[0].ID(),
				CreatedAt: info[0].CreatedAt(),
				UpdatedAt: info[0].UpdatedAt(),
			}),
		)

	res, err := s.repo.GetAddStat(time.Now(), time.Now())

	require.NoError(s.T(), err)
	require.Equal(s.T(), info, res)
}

func TestGetAddStatSuite(t *testing.T) {
	suite.Run(t, new(GetAddStatSuite))
}
