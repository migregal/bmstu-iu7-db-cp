//go:build smoke
// +build smoke

package database

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupTest() {
}

func (s *TestSuite) TestPostgreSQL() {
	interactor, err := New(Params{
		Host:     "localhost",
		Port:     "5432",
		User:     "user",
		Password: "password",
		DBName:   "smoke_db",
		Driver:   "pg",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), interactor)

	var temp string
	interactor.Raw("SELECT version()").Scan(&temp)

	require.True(s.T(), regexp.MustCompile(`^PostgreSQL 14.2.*`).MatchString(temp))
}

func (s *TestSuite) TestMySQL() {
	interactor, err := New(Params{
		Host:     "localhost",
		Port:     "3306",
		User:     "user",
		Password: "password",
		DBName:   "smoke_db",
		Driver:   "mysql",
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), interactor)

	var temp string
	interactor.Raw("SELECT VERSION()").Scan(&temp)
	require.True(s.T(), regexp.MustCompile(`^8.0.*`).MatchString(temp))
}

func (s *TestSuite) TearDownTest() {
}

func TestRegisterSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
