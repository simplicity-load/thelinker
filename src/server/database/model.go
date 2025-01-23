package database

import "time"

type Shortlink struct {
	OriginalURL string    `json:"original_url" db:"original_url"`
	ShortURL    string    `json:"short_url" db:"short_url"`
	Date        time.Time `json:"date" db:"date"`
}
