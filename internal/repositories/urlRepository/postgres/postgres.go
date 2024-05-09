package postgres

import (
	"context"
	"github.com/vinylSummer/microUrl/pkg/postgres"
)

type URLRepository struct {
	*postgres.Connection
}

func New(db *postgres.Connection) *URLRepository {
	return &URLRepository{db}
}

func (repo *URLRepository) StoreURLsBinding(longURL string, shortURL string, ctx context.Context) error {
	return nil
}

func (repo *URLRepository) GetLongURL(shortURL string, ctx context.Context) (string, error) {
	return "", nil
}

func (repo *URLRepository) CheckUnique(shortURL string) (bool, error) {
	return false, nil
}
