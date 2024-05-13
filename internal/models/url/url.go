package models

import "time"

type ShortURL struct {
	Value string `json:"value"`
}

type LongURL struct {
	Value string `json:"value"`
}

type URLBinding struct {
	ID        uint      `json:"id"`
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at" example:"2024-05-09T11:05:43+09:00"`
}
