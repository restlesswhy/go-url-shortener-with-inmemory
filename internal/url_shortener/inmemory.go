package urlshortener

type UrlShortenerInmemory interface {
	GetShortInmemory(longUrl string) (string, error) 
	CreateInmemory(shortUrl, longUrl string) error
	GetLongInmemory(shortUrl string) (string, error) 

}