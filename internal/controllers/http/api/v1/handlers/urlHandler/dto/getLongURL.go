package dto

import models "github.com/vinylSummer/microUrl/internal/models/url"

type GetLongURLRequest struct {
	ShortURL string `json:"short_url"`
}

type GetLongURLResponse struct {
	LongURL string `json:"long_url"`
}

func (request *GetLongURLRequest) ToModel() *models.ShortURL {
	return &models.ShortURL{
		Value: request.ShortURL,
	}
}
