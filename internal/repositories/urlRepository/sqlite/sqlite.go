package sqlite

import (
	"context"
	"fmt"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"time"
)

type URLRepository struct {
	*sqlite.Connection
}

func New(db *sqlite.Connection) *URLRepository {
	return &URLRepository{db}
}

func (repo *URLRepository) StoreURLsBinding(longURL string, shortURL string, createdAt time.Time, ctx context.Context) error {
	fmt.Println("SQLiteURLRepository.StoreURLsBinding...")
	return nil
}

func (repo *URLRepository) GetLongURL(shortURL string, ctx context.Context) (string, error) {
	fmt.Println("SQLiteURLRepository.StoreURLsBinding...")
	return "nil", nil
}
