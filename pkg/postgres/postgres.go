package postgres

import (
	"github.com/jackc/pgx"
)

type Connection struct {
	DB *pgx.Conn
}

func New(url string) (*Connection, error) {
	connectionConfig, err := pgx.ParseConnectionString(url)
	if err != nil {
		return nil, err
	}

	connection, err := pgx.Connect(connectionConfig)
	if err != nil {
		return nil, err
	}

	return &Connection{DB: connection}, nil
}

func (postgresConnection *Connection) Close() {
	if postgresConnection.DB != nil {
		postgresConnection.DB.Close()
	}
}
