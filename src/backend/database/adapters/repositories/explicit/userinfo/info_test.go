//go:build unit
// +build unit

package userinfo

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"neural_storage/database/core/services/interactor/database"
)

type TestSuite struct {
	suite.Suite

	SqlDB   *sql.DB
	DB      *gorm.DB
	SqlMock sqlmock.Sqlmock

	repo Repository
}

func (s *TestSuite) SetupTest() {
	var err error
	s.SqlDB, s.SqlMock, err = sqlmock.New()
	require.NoError(s.T(), err)
	require.NotNil(s.T(), s.SqlDB)
	require.NotNil(s.T(), s.SqlMock)

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 s.SqlDB,
		PreferSimpleProtocol: true,
	})

	s.DB, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(s.T(), err)

	s.repo = Repository{db: database.Interactor{DB: s.DB}}
}

func (s *TestSuite) TearDownTest() {
	s.SqlDB.Close()
}
