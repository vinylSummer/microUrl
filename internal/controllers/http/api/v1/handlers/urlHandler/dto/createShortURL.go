package dto

type CreateShortURLRequest struct {
	LongURL string `json:"long_url"`
}

type CreateShortURLResponse struct {
	ShortURL string `json:"short_url"`
}

//func (request *CreateShortURLRequest) ToModel() *models.LongURL {
//	return &models.LongURL{
//		Value: request.LongURL,
//	}
//}
