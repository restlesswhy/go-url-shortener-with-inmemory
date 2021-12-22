package inmemory

import (
	"testing"

	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
	inmemory "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/memdb"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"github.com/stretchr/testify/require"
)


func TestUrlShortenerInmemory_CreateInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUSInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	errCreate := shortMemDB.Create(model.ShortUrl, model.LongUrl)

	shortUrl, _ := shortMemDB.GetShort(model.LongUrl)

	require.Nil(t, errCreate)
	require.Equal(t, model.ShortUrl, shortUrl)
}

func TestUrlShortenerInmemory_GetShortInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUSInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	_ = shortMemDB.Create(model.ShortUrl, model.LongUrl)

	shortUrl, errGet := shortMemDB.GetShort(model.LongUrl)

	require.Nil(t, errGet)
	require.Equal(t, model.ShortUrl, shortUrl)
}

func TestUrlShortenerInmemory_GetShortInmemoryHaveNo(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUSInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	otherShortUrl := "aasdasdasd"

	_ = shortMemDB.Create(model.ShortUrl, model.LongUrl)

	shortUrl, errGet := shortMemDB.GetShort(otherShortUrl)

	logger.Info(shortUrl, errGet)
	require.Empty(t, shortUrl)
}

func TestUrlShortenerInmemory_GetLongInmemory(t *testing.T) {
	memdb, _ := inmemory.InitMemDB()

	shortMemDB := NewUSInmemory(memdb)

	model := models.UrlsLS{
		ShortUrl: "some short",
		LongUrl: "some long",
	}

	_ = shortMemDB.Create(model.ShortUrl, model.LongUrl)

	longUrl, errGet := shortMemDB.GetLong(model.ShortUrl)

	require.Nil(t, errGet)
	require.Equal(t, model.LongUrl, longUrl)
}