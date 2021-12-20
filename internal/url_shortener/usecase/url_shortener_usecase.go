package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"github.com/speps/go-hashids"
)

// UrlShortenerUC image useCase
type UrlShortenerUC struct {
	shortenerRepo us.UrlShortenerRepository
	cfg *config.Config
	memdb us.UrlShortenerInmemory
}

// NewUrlShortenerUC image useCase constructor
func NewUrlShortenerUC(cfg *config.Config, shortenerRepo us.UrlShortenerRepository, memdb us.UrlShortenerInmemory) *UrlShortenerUC {
	return &UrlShortenerUC{
		shortenerRepo: shortenerRepo,
		cfg: cfg,
		memdb: memdb,
	}
}

// Create is create new short url
func (u *UrlShortenerUC) Create(ctx context.Context, longUrl string) (string, error) {
	if longUrl == "" {
		logger.Error("empty string in UrlShortenerUC.create")
		return "your url is empty", nil
	}

	logger.Infof("New creating with long url===============>%s", longUrl)
	
	shortUrl, err := u.memdb.GetShortInmemory(longUrl) // Проверяем есть ли урл в локальном хранилище 
	if err != nil {
		return shortUrl, err
	}

	if shortUrl == "" { // Если локально урл не найден, создаем для метода GetRepo и ищем в нем
		shortUrl, err = u.getUniqueString(longUrl)
		if err != nil {
			return shortUrl, err
		}

		urls, ok := u.shortenerRepo.GetRepo(ctx, longUrl, shortUrl) 
		if ok { // Если урл найден в базе, сохраняем его локально и возвращаем найденный урл
			if err := u.memdb.CreateInmemory(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
			return urls.ShortUrl, nil
		}

		// logger.Infof("urls in model: %s, %s", urls.LongUrl, urls.ShortUrl)
		if !ok { // Если урл не найден в базе, создаем урл в базе и локально, и возвращаем созданный урл если все прошло успешно
			if err := u.shortenerRepo.CreateRepo(ctx, longUrl, shortUrl); err != nil {
				return "", errors.Wrap(err, "u.shortenerRepo.CreateRepo")
			}
			if err := u.memdb.CreateInmemory(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
			return shortUrl, nil
		}
	}
	return shortUrl, nil
	// return "", errors.Wrap(errors.New("something went wrong"), "u.shortenerUC.Create")
}

// Get return long url
func (u *UrlShortenerUC) Get(ctx context.Context, shortUrl string) (string, error) {
	if shortUrl == "" {
		logger.Error("empty string in UrlShortenerUC.get")
		return "your url is empty", nil
	}

	logger.Infof("New getting with short url===============>%s", shortUrl)

	longUrl, err := u.memdb.GetLongInmemory(shortUrl) // Проверяем есть ли урл локально
	if err != nil {
		return shortUrl, err
	}

	if longUrl == "" { // Если локально длинный урл не найден ищем в базе 
		urls, ok := u.shortenerRepo.GetRepo(ctx, longUrl, shortUrl) 
		if ok { // Если урл найден в базе, создаем его локально и возвращаем
			if err := u.memdb.CreateInmemory(urls.ShortUrl, urls.LongUrl); err != nil { 
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}
			return urls.LongUrl, nil
		}

		if !ok { // Если в базе урл не найден, сообщаем об этом
			return "this short url is not exist", nil
		}
		longUrl = urls.LongUrl
	}
	return longUrl, nil
}

// getUniqueString create unique short url
func (u *UrlShortenerUC) getUniqueString(longUrl string) (string, error) {
	hd := hashids.NewData()
	
	hd.Salt = longUrl
	hd.Alphabet = u.cfg.Shortener.Runes
	hd.MinLength = u.cfg.Shortener.StringLength

	h, err := hashids.NewWithData(hd)
	if err != nil {
		logger.Error("cant encode string")
		return "", err
	}

	e, err := h.Encode([]int{45, 434, 1313, 99})
	if err != nil {
		logger.Error("cant encode string")
		return "", err
	}

	return e, nil
}

