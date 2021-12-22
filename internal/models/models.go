package models

import "time"

type UrlsLS struct {
	LongUrl  string `db:"long_url"`
	ShortUrl string `db:"short_url"`
	ExpiresAt time.Time
}
