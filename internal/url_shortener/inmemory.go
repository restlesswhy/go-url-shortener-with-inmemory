package urlshortener

type UrlShortenerInmemory interface {
	GetLongInmemory(longUrl string) (string, error) 
	CreateInmemory(shortUrl, longUrl string) error
}