package sqlite

import (
	ctx "context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"time"
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
		log.Error().Err(err).Msg("Couldn't begin transaction")
		return err
	}

	statement, err := transaction.Prepare("INSERT INTO url_bindings (short_url, long_url) VALUES (?, ?)")
	defer statement.Close()
	if err != nil {
		log.Error().Err(err).Msg("Couldn't prepare statement")
		return err
	}

	// TODO: if unique constraint fails - return custom err with existing short URL
	_, err = statement.Exec(binding.ShortURL, binding.LongURL)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return err
	}

	err = transaction.Commit()
	if err != nil {
		log.Error().Err(err).Msg("Couldn't commit transaction")
		return err
	}

	return nil
}

func (repo *URLRepository) GetURLsBinding(context ctx.Context, shortURL string) (*models.URLBinding, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT id, long_url, created_at FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Error().Err(err).Msg("Couldn't prepare statement")
		return nil, err
	}
	defer statement.Close()

	var id uint
	var longURL string
	var createdAtString string
	err = statement.QueryRowContext(context, shortURL).Scan(&id, &longURL, &createdAtString)
	if errors.Is(err, sql.ErrNoRows) {
		log.Error().Err(err).Msgf("Couldn't find long URL for %s in the database", shortURL)
		return nil, &ErrURLNotFound{URL: shortURL}
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, createdAtString)

	log.Info().Msgf("Retrieved the url binding for short URL: %s", shortURL)

	return &models.URLBinding{
		ID:        id,
		LongURL:   longURL,
		ShortURL:  shortURL,
		CreatedAt: createdAt,
	}, nil
}

func (repo *URLRepository) GetLongURL(context ctx.Context, shortURL string) (string, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT long_url FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Error().Err(err).Msg("Couldn't prepare statement")
		return "", err
	}
	defer statement.Close()

	var longURL string
	err = statement.QueryRowContext(context, shortURL).Scan(&longURL)
	if errors.Is(err, sql.ErrNoRows) {
		log.Error().Err(err).Msgf("Couldn't find long URL for %s in the database", shortURL)
		return "", &ErrURLNotFound{URL: shortURL}
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return "", err
	}

	log.Info().Msgf("Retrieved long URL: %s for short URL: %s", longURL, shortURL)

	return longURL, nil
}

func (repo *URLRepository) GetShortURL(context ctx.Context, longURL string) (string, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT short_url FROM url_bindings WHERE long_url = ?")
	if err != nil {
		log.Error().Err(err).Msg("Couldn't prepare statement")
		return "", err
	}
	defer statement.Close()

	var shortURL string
	err = statement.QueryRowContext(context, longURL).Scan(&shortURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", &ErrURLNotFound{URL: shortURL}
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return "", err
	}

	log.Info().Msgf("Retrieved short URL: %s for long URL: %s", shortURL, longURL)

	return shortURL, nil
}

func (repo *URLRepository) CheckUnique(context ctx.Context, shortURL string) (bool, error) {
	statement, err := repo.Connection.DB.Prepare("SELECT id FROM url_bindings WHERE short_url = ?")
	if err != nil {
		log.Error().Err(err).Msg("Couldn't prepare statement")
		return false, err
	}
	defer statement.Close()

	var id uint
	err = statement.QueryRowContext(context, shortURL).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return false, err
	}
	return false, nil
}
