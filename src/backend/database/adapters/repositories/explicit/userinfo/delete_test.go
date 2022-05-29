//go:build unit
// +build unit

package userinfo

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"neural_storage/cube/core/entities/user"
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
	id := "test"
	expected := *user.NewInfo(id, "", "", "", "", 0, time.Time{})

	s.SqlMock.
		ExpectExec(`^DELETE FROM "users_info" WHERE "users_info"."id"`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.repo.Delete(expected)

	require.NoError(s.T(), err)
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}
