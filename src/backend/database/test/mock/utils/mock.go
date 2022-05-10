package utils

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm/schema"
)

func MockRows(objs ...interface{}) *sqlmock.Rows {
	if reflect.ValueOf(objs[0]).Kind() == reflect.Slice {
		if len(objs) > 1 {
			panic(fmt.Errorf("objs must have one element if first element is slice"))
		}

		s := reflect.ValueOf(objs[0])
		if s.IsNil() {
			return nil
		}

		objs = make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			objs[i] = s.Index(i).Interface()
		}
	}

	s, err := schema.Parse(objs[0], &sync.Map{}, schema.NamingStrategy{})
	if err != nil {
		panic("failed to create schema")
	}

	columns := make([]string, 0)
	for _, field := range s.Fields {
		if len(field.DBName) == 0 {
			continue
		}
		columns = append(columns, field.DBName)
	}

	rows := sqlmock.NewRows(columns)

	for _, obj := range objs {
		row := make([]driver.Value, 0)

		for _, field := range s.Fields {
			if len(field.DBName) == 0 {
				continue
			}
			r := reflect.ValueOf(obj)
			f := reflect.Indirect(r).FieldByName(field.Name)
			row = append(row, f.Interface())
		}

		rows = rows.AddRow(row...)
	}

	return rows
}
