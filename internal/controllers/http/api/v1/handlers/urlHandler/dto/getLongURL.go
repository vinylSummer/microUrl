package dto

type GetLongURLRequest struct {
	ShortURL string `json:"short_url"`
}

type GetLongURLResponse struct {
	LongURL string `json:"long_url"`
}
