package services

import (
	"context"
	"github.com/vinylSummer/microUrl/internal/repositories"
	"time"
)

type URLService struct {
	urlRepo repositories.URLRepository
}

func NewURLService(urlRepo repositories.URLRepository) *URLService {
	return &URLService{urlRepo: urlRepo}
}

func (service *URLService) generateShortURL(longURL string) (string, error) {
	shortURL := "benis.haha/hehe"

	return shortURL, nil
}

func (service *URLService) GetLongURL(shortURL string) (string, error) {
	longURL, err := service.urlRepo.GetLongURL(shortURL, context.Background())
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (service *URLService) CreateShortURL(longURL string) (string, error) {
	shortURL, err := service.generateShortURL(longURL)
	if err != nil {
		return "", err
	}

	err = service.urlRepo.StoreURLsBinding(longURL, shortURL, time.Now(), context.Background())
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
