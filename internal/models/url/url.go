package models

import "time"

// I'm not THAT sure yet if i should use DTOs as service returns yet
//
//	type URLBinding struct {
//		LongURL   LongURL    `json:"long_url" example:"https://www.instagram.com/p/C6vSstIyYsC/?hl=en&img_index=1"`
//		ShortURL  ShortURL    `json:"short_url" example:"murl.xyz/genius"`
//		CreatedAt time.Time `json:"created_at" example:"2024-05-09T11:05:43+09:00"`
//	}
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
