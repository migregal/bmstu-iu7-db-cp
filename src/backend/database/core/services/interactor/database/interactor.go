package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Interactor struct {
	*gorm.DB
}

type Params struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Driver   string
}

func New(p Params) (Interactor, error) {
	conn, err := gorm.Open(newDBDialector(p), &gorm.Config{QueryFields: true})
	if err != nil {
		return Interactor{}, err
	}

	return Interactor{DB: conn}, nil
}

func newDBDialector(p Params) gorm.Dialector {
	switch p.Driver {
	case "pg":
		dbConfig := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			p.Host,
			p.Port,
			p.DBName,
			p.User,
			p.Password,
			"disable",
		)

		return postgres.Open(dbConfig)
	case "mysql":
		dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			p.User,
			p.Password,
			p.Host,
			p.Port,
			p.DBName,
		)

		return mysql.Open(dbConfig)
	}

	panic("wrong db driver specified")
}
