package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Connection struct {
	DB *sql.DB
}

type DatabaseDoesNotExistError struct {
	err string
	url string
}

func (error *DatabaseDoesNotExistError) Error() string {
	return fmt.Sprintf("Couldn't connect to database at %s, %s", error.url, error.err)
}

func New(url string) (*Connection, error) {
	if _, err := os.Stat(url); errors.Is(err, os.ErrNotExist) {
		return nil, &DatabaseDoesNotExistError{url: url, err: err.Error()}
	}
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
