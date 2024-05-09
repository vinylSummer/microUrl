package repositories

import "context"
import "time"

type URLRepository interface {
	StoreURLsBinding(longURL string, shortURL string, createdAt time.Time, ctx context.Context) error
	GetLongURL(shortURL string, ctx context.Context) (string, error)
}
