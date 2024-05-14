package sqlite

import (
	ctx "context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
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
		log.Error().Err(err).Msg(fmt.Sprintf("Couldn't find long URL for %s in the database", shortURL))
		return "", &ErrURLNotFound{URL: shortURL}
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return "", err
	}

	log.Info().Msg(fmt.Sprintf("Retrieved long URL: %s for short URL: %s", longURL, shortURL))

	return longURL, nil
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
		log.Error().Err(err).Msg(fmt.Sprintf("Couldn't find long URL for %s in the database", shortURL))
		return true, nil
	}
	if err != nil {
		log.Error().Err(err).Msg("Couldn't execute statement")
		return false, err
	}
	return false, nil
}
