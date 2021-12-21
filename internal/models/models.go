package models

import "time"

type UrlsLS struct {
	LongUrl  string `db:"long_url"`
	ShortUrl string `db:"short_url"`
	CreateAt  time.Time `db:"create_at"`
	ExpiresAt time.Time `db:"expires_at"`
}
