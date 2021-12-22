package inmemory

import (
	"testing"

	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
	inmemory "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/memdb"
	"github.com/stretchr/testify/require"
)


func TestUrlShortenerInmemory_CreateInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUrlShortenerInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	errCreate := shortMemDB.CreateInmemory(model.ShortUrl, model.LongUrl)

	shortUrl, _ := shortMemDB.GetShortInmemory(model.LongUrl)

	require.Nil(t, errCreate)
	require.Equal(t, model.ShortUrl, shortUrl)
}

func TestUrlShortenerInmemory_GetShortInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUrlShortenerInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	_ = shortMemDB.CreateInmemory(model.ShortUrl, model.LongUrl)

	shortUrl, errGet := shortMemDB.GetShortInmemory(model.LongUrl)

	require.Nil(t, errGet)
	require.Equal(t, model.ShortUrl, shortUrl)
}

func TestUrlShortenerInmemory_GetLongInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUrlShortenerInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	_ = shortMemDB.CreateInmemory(model.ShortUrl, model.LongUrl)

	longUrl, errGet := shortMemDB.GetLongInmemory(model.ShortUrl)

	require.Nil(t, errGet)
	require.Equal(t, model.LongUrl, longUrl)
}