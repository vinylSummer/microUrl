package v1

import (
	ctx "context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/internal/repositories"
	"github.com/vinylSummer/microUrl/internal/repositories/urlRepository/sqlite"
	"math/rand/v2"
	"time"
)

type URLService struct {
	urlRepo   repositories.URLRepository
	cacheRepo repositories.CacheRepository
}

func NewURLService(urlRepo repositories.URLRepository, cacheRepo repositories.CacheRepository) *URLService {
	return &URLService{
		urlRepo:   urlRepo,
		cacheRepo: cacheRepo,
	}
}

// TODO: rewrite for better scaling
func (service *URLService) generateUniqueString(context ctx.Context, length int) (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	maxRetries := 5
	for range maxRetries {
		randomCharacters := make([]rune, length)
		for i := range randomCharacters {
			randomCharacters[i] = letters[rand.IntN(len(letters))]
		}
		randomCharactersString := string(randomCharacters)

		isUnique, err := service.urlRepo.CheckUnique(context, randomCharactersString)
		if err != nil {
			log.Error().Err(err).Msgf("Error while checking if URL %s is unique", randomCharactersString)
			continue
		}
		if isUnique {
			return randomCharactersString, nil
		}
	}
	return "", fmt.Errorf("couldn't generate unique string in %s retries", maxRetries)
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (service *URLService) generateShortURL(context ctx.Context) (string, error) {
	URLLength := randRange(3, 7)

	shortURL, err := service.generateUniqueString(context, URLLength)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (service *URLService) GetLongURL(context ctx.Context, shortURL *models.ShortURL) (string, error) {
	longURL, err := service.getCachedLongURL(context, shortURL)
	if err != nil || longURL == "" {
		longURL, err = service.urlRepo.GetLongURL(context, shortURL.Value)
		if err != nil {
			return "", err
		}

		urlsBinding, err := service.urlRepo.GetURLsBinding(context, shortURL.Value)
		if err != nil {
			log.Error().Err(err).Msgf("Error while fetching urls binding for short URL %s", shortURL)
		}

		err = service.cacheURLBinding(context, urlsBinding)
		if err != nil {
			log.Error().Err(err).Msgf("Error while trying to cache URLs binding: %s -> %s", shortURL, urlsBinding.LongURL)
		}
	}

	return longURL, nil
}

func (service *URLService) getCachedLongURL(context ctx.Context, shortURL *models.ShortURL) (string, error) {
	longURL, err := service.cacheRepo.Get(context, shortURL.Value)
	if err != nil {
		return "", err
	}

	longURLString, ok := longURL.(string)
	if !ok {
		return "", errors.New("cached longURL is not a string")
	}

	log.Info().Msgf("Retrieved longURL: %s from cache", longURLString)

	return longURLString, nil
}

func (service *URLService) CreateShortURL(context ctx.Context, longURL *models.LongURL) (*models.URLBinding, error) {
	shortURL, err := service.urlRepo.GetShortURL(context, longURL.Value)
	var urlNotFoundErr *sqlite.ErrURLNotFound
	if err != nil && !errors.As(err, &urlNotFoundErr) {
		return nil, err
	}
	if shortURL != "" {
		urlBinding := &models.URLBinding{
			ShortURL: shortURL,
			LongURL:  longURL.Value,
		}

		return urlBinding, nil
	}

	urlBinding := &models.URLBinding{
		ShortURL: shortURL,
		LongURL:  longURL.Value,
	}

	err = service.urlRepo.StoreURLsBinding(context, urlBinding)
	if err != nil {
		return nil, err
	}

	return urlBinding, nil
}

func (service *URLService) cacheURLBinding(context ctx.Context, binding *models.URLBinding) error {
	err := service.cacheRepo.Set(context, binding.ShortURL, binding.LongURL, time.Hour*1)
	if err != nil {
		return err
	}
	return nil
}
