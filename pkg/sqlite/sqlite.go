package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

type Connection struct {
	DB *sql.DB
}

func createTables(db *sql.DB) {
	createURLBindingsTable :=
		`CREATE TABLE IF NOT EXISTS url_bindings (
			id INTEGER PRIMARY KEY,
			short_url TEXT NOT NULL UNIQUE,
			long_url TEXT NOT NULL,
			created_at TEXT DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := db.Exec(createURLBindingsTable)
	if err != nil {
		log.Fatal().Err(err)
	}

}

func New(url string) (*Connection, error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}

	createTables(db)

	return &Connection{DB: db}, nil
}

func (sqliteConnection *Connection) Close() {
	if sqliteConnection.DB != nil {
		sqliteConnection.DB.Close()
	}
}
