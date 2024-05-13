package dto

type GetLongURLRequest struct {
	ShortURL string `json:"shortURL"`
}

type GetLongURLResponse struct {
	LongURL string `json:"longURL"`
}
