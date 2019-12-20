package mocks

import (
	"database/sql"
	"database/sql/driver"
	"github.com/stretchr/testify/mock"
)

type dbMock struct {
	mock.Mock
}

func CreateDBMock() *dbMock {
	return &dbMock{}
}

func (m *dbMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return &sql.Rows{}, nil
}

func (m *dbMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	return driver.ResultNoRows, nil
}

func (m *dbMock) QueryRow(query string, args ...interface{}) *sql.Row {
	return &sql.Row{}
}
