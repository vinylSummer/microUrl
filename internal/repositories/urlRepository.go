package repositories

import "context"

type URLRepository interface {
	StoreURLsBinding(longURL string, shortURL string, ctx context.Context) error
	GetLongURL(shortURL string, ctx context.Context) (string, error)
	CheckUnique(shortURL string) (bool, error)
}
