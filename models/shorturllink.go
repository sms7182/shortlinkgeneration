package models

import "time"

type ShortUrlLink struct {
	ID        int64     `json:"id"`
	Url       string    `json:"url"`
	ShortLink string    `json:"short_link"`
	Count     int64     `json:"count"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
