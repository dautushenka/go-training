package core

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file:data.db?cache=shared&mode=memory")
	if err != nil {
		panic(err)
	}

	createSchema(db)
	applyFixtures(db)

	return db
}
