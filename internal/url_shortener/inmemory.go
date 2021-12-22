//go:generate mockgen -source inmemory.go -destination mock/inmemory.go -package mock
package urlshortener

type USInmemory interface {
	GetShort(longUrl string) (string, error) 
	Create(shortUrl, longUrl string) error
	GetLong(shortUrl string) (string, error) 
	Check() error
}