package services

import (
	ctx "context"
	models "github.com/vinylSummer/microUrl/internal/models/url"
)

type URLService interface {
	GetLongURL(context ctx.Context, shortURL *models.ShortURL) (string, error)
	CreateShortURL(context ctx.Context, longURL *models.LongURL) (*models.URLBinding, error)
}
