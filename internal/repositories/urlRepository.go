package repositories

import (
	ctx "context"
	models "github.com/vinylSummer/microUrl/internal/models/url"
)

type URLRepository interface {
	StoreURLsBinding(context ctx.Context, binding *models.URLBinding) error
	GetURLsBinding(context ctx.Context, shortURL string) (*models.URLBinding, error)
	GetLongURL(context ctx.Context, shortURL string) (string, error)
	GetShortURL(context ctx.Context, longURL string) (string, error)
	CheckUnique(context ctx.Context, shortURL string) (bool, error)
}
