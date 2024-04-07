package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", ("file:" + path))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
