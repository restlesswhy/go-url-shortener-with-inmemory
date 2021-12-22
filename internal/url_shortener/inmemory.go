//go:generate mockgen -source inmemory.go -destination mock/inmemory.go -package mock
package urlshortener

type UrlShortenerInmemory interface {
	GetShortInmemory(longUrl string) (string, error) 
	CreateInmemory(shortUrl, longUrl string) error
	GetLongInmemory(shortUrl string) (string, error) 
	CheckInmemory()
}