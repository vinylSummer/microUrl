package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"log"
)

type URLRepository struct {
	*sqlite.Connection
}

func New(db *sqlite.Connection) *URLRepository {
	return &URLRepository{db}
}

func (repo *URLRepository) StoreURLsBinding(longURL string, shortURL string, ctx context.Context) error {
	transaction, err := repo.Connection.DB.Begin()
	if err != nil {
		log.Printf("Couldn't begin transaction in StoreURLsBinding: %v", err)
		return err
	}

	statement, err := transaction.Prepare("INSERT INTO url_bindings (short_url, long_url) VALUES (?, ?)")
	if err != nil {
		log.Printf("Couldn't prepare statement in StoreURLsBinding: %v", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(shortURL, longURL)
	if err != nil {
		log.Printf("Couldn't execute statement in StoreURLsBinding: %v", err)
		return err
	}

	err = transaction.Commit()
	if err != nil {
		log.Printf("Couldn't commit transaction in StoreURLsBinding: %v", err)
		return err
	}

	return nil
}

func (repo *URLRepository) GetLongURL(shortURL string, ctx context.Context) (string, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT long_url FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Printf("Couldn't prepare statement in GetLongURL: %v", err)
		return "", err
	}
	defer statement.Close()

	var longURL string
	err = statement.QueryRow(shortURL).Scan(&longURL)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("Couldn't find long URL for %s in the database: %v", shortURL, err)
		return "", sql.ErrNoRows
	}
	if err != nil {
		log.Printf("Couldn't execute statement in GetLongURL: %v", err)
		return "", err
	}
	log.Printf("Retrieved long URL: %s for short URL: %s", longURL, shortURL)

	return longURL, nil
}

func (repo *URLRepository) CheckUnique(shortURL string) (bool, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT id FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Printf("Couldn't prepare statement in GetLongURL: %v", err)
		return false, err
	}
	defer statement.Close()

	var id uint
	err = statement.QueryRow(shortURL).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("Couldn't find long URL for %s in the database: %v", shortURL, err)
		return true, nil
	}
	if err != nil {
		log.Printf("Couldn't execute statement in GetLongURL: %v", err)
		return false, err
	}
	return false, nil
}
