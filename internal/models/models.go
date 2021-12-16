package models

type UrlsLS struct {
	LongUrl  string `db:"long_url"`
	ShortUrl string `db:"short_url"`
}
