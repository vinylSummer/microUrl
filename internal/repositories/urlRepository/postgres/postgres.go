package postgres

import (
	ctx "context"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/pkg/postgres"
)

type URLRepository struct {
	*postgres.Connection
}

func New(db *postgres.Connection) *URLRepository {
	return &URLRepository{db}
}

func (repo *URLRepository) StoreURLsBinding(context ctx.Context, binding *models.URLBinding) error {
	return nil
}

func (repo *URLRepository) GetURLsBinding(context ctx.Context, shortURL string) (*models.URLBinding, error) {
	return nil, nil
}

func (repo *URLRepository) GetLongURL(context ctx.Context, shortURL string) (string, error) {
	return "", nil
}

func (repo *URLRepository) GetShortURL(context ctx.Context, longURL string) (string, error) {
	return "", nil
}

func (repo *URLRepository) CheckUnique(context ctx.Context, shortURL string) (bool, error) {
	return false, nil
}
