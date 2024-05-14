package sqlite

import (
	ctx "context"
	"database/sql"
	"errors"
	"fmt"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"log"
)

type ErrURLNotFound struct {
	URL string
}

func (e ErrURLNotFound) Error() string {
	return fmt.Sprintf("URL %s not found", e.URL)
}

type URLRepository struct {
	*sqlite.Connection
}

func New(db *sqlite.Connection) *URLRepository {
	return &URLRepository{db}
}

func (repo *URLRepository) StoreURLsBinding(context ctx.Context, binding *models.URLBinding) error {
	transaction, err := repo.Connection.DB.BeginTx(context, nil)
	defer transaction.Rollback()
	if err != nil {
		log.Printf("Couldn't begin transaction in StoreURLsBinding: %v", err)
		return err
	}

	statement, err := transaction.Prepare("INSERT INTO url_bindings (short_url, long_url) VALUES (?, ?)")
	defer statement.Close()
	if err != nil {
		log.Printf("Couldn't prepare statement in StoreURLsBinding: %v", err)
		return err
	}

	_, err = statement.Exec(binding.ShortURL, binding.LongURL)
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

func (repo *URLRepository) GetLongURL(context ctx.Context, shortURL string) (string, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT long_url FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Printf("Couldn't prepare statement in GetLongURL: %v", err)
		return "", err
	}
	defer statement.Close()

	var longURL string
	err = statement.QueryRowContext(context, shortURL).Scan(&longURL)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("Couldn't find long URL for %s in the database: %v", shortURL, err)
		return "", &ErrURLNotFound{URL: shortURL}
	}
	if err != nil {
		log.Printf("Couldn't execute statement in GetLongURL: %v", err)
		return "", err
	}
	log.Printf("Retrieved long URL: %s for short URL: %s", longURL, shortURL)

	return longURL, nil
}

func (repo *URLRepository) CheckUnique(context ctx.Context, shortURL string) (bool, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT id FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Printf("Couldn't prepare statement in GetLongURL: %v", err)
		return false, err
	}
	defer statement.Close()

	var id uint
	err = statement.QueryRowContext(context, shortURL).Scan(&id)
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
