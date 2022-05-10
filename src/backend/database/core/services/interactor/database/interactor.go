package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Interactor struct {
	*gorm.DB
}

type Params struct {
	Host     string
	User     string
	Password string
	DBName   string
	Driver   string
}

func New(p Params) (Interactor, error) {
	conn, err := gorm.Open(newDBDialector(p), &gorm.Config{})
	if err != nil {
		return Interactor{}, err
	}

	return Interactor{DB: conn}, nil
}

func newDBDialector(p Params) gorm.Dialector {
	switch p.Driver {
	case "pg":
		dbConfig := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
			p.Host,
			p.User,
			p.DBName,
			p.Password,
			"disable",
		)

		return postgres.Open(dbConfig)
	}

	panic("wrong db driver specified")
}
