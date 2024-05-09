package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	DB *sql.DB
}

func New(url string) (*Connection, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}

	return &Connection{DB: db}, nil
}

func (sqliteConnection *Connection) Close() {
	if sqliteConnection.DB != nil {
		sqliteConnection.DB.Close()
	}
}
