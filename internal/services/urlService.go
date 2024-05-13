package services

import (
	"context"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/internal/repositories"
	"log"
	"math/rand/v2"
)

type URLService struct {
	urlRepo repositories.URLRepository
}

func NewURLService(urlRepo repositories.URLRepository) *URLService {
	return &URLService{
		urlRepo: urlRepo,
	}
}

func (service *URLService) generateUniqueString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for {
		randomCharacters := make([]rune, length)
		for i := range randomCharacters {
			randomCharacters[i] = letters[rand.IntN(len(letters))]
		}
		randomCharactersString := string(randomCharacters)

		isUnique, err := service.urlRepo.CheckUnique(randomCharactersString)
		if err != nil {
			log.Printf("Error while checking if url is unique: %s", err)
			continue
		}
		if isUnique {
			return randomCharactersString
		}
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (service *URLService) generateShortURL() (string, error) {
	URLLength := randRange(1, 7)

	shortURL := service.generateUniqueString(URLLength)

	return shortURL, nil
}

func (service *URLService) GetLongURL(shortURL *models.ShortURL) (string, error) {
	longURL, err := service.urlRepo.GetLongURL(shortURL.Value, context.Background())
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (service *URLService) CreateShortURL(longURL *models.LongURL) (string, error) {
	shortURL, err := service.generateShortURL()
	if err != nil {
		return "", err
	}

	err = service.urlRepo.StoreURLsBinding(longURL.Value, shortURL, context.Background())
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
