package v1

import (
	ctx "context"
	"github.com/rs/zerolog/log"
	models "github.com/vinylSummer/microUrl/internal/models/url"
	"github.com/vinylSummer/microUrl/internal/repositories"
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

// TODO: rewrite for better scaling
func (service *URLService) generateUniqueString(context ctx.Context, length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for {
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
			return randomCharactersString
		}
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (service *URLService) generateShortURL(context ctx.Context) (string, error) {
	URLLength := randRange(1, 7)

	shortURL := service.generateUniqueString(context, URLLength)

	return shortURL, nil
}

func (service *URLService) GetLongURL(context ctx.Context, shortURL *models.ShortURL) (string, error) {
	longURL, err := service.urlRepo.GetLongURL(context, shortURL.Value)
	if err != nil {
		return "", err
	}

	return longURL, nil
}

func (service *URLService) CreateShortURL(context ctx.Context, longURL *models.LongURL) (*models.URLBinding, error) {
	shortURL, err := service.generateShortURL(context)
	if err != nil {
		return nil, err
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
