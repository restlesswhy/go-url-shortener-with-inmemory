package usecase

import (
	"context"
	"crypto/sha512"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	us "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
	"github.com/speps/go-hashids"
)

// USUseCase image useCase
type USUseCase struct {
	shortenerRepo us.USRepository
	cfg           *config.Config
	memdb         us.USInmemory
}

type single struct {
}

var singleInstance *single

// NewUSUseCase image useCase constructor
func NewUSUseCase(cfg *config.Config, shortenerRepo us.USRepository, memdb us.USInmemory) *USUseCase {
	return &USUseCase{
		shortenerRepo: shortenerRepo,
		cfg:           cfg,
		memdb:         memdb,
	}
}

// Create is create new short url
func (u *USUseCase) Create(ctx context.Context, longUrl string) (string, error) {
	// Запускаем проверку времени урлов в inmemory через синглтон
	var once sync.Once
	if singleInstance == nil {
        once.Do(
            func() {
				u.CallAt(0, 0, 0, time.Hour, u.memdb.Check)
                singleInstance = &single{}
            })
    }

	logger.Infof("New creating with long url===============>%s", longUrl)

	if longUrl == "" {
		return "Your url is empty", nil
	}
	// Проверяем есть ли урл в локальном хранилище
	shortUrl, err := u.memdb.GetShort(longUrl) 
	if err != nil {
		return shortUrl, errors.Wrap(err, "memdb.GetShort")
	}

	if shortUrl == "" { 
		// Если локально урл не найден, ищем в базе
		urls, ok, err := u.shortenerRepo.Get(ctx, longUrl, "")
		if err != nil {
			return "", errors.Wrap(err, "shortenerRepo.Get")
		}

		switch ok {
		// Если урл найден в базе, сохраняем его локально и возвращаем найденный урл
		case true:
			if err := u.memdb.Create(urls.ShortUrl, urls.LongUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}

			logger.Infof("Return short url - %s", urls.ShortUrl)
			return urls.ShortUrl, nil
		// Если урл не найден в базе, создаем урл в базе и локально, и возвращаем созданный урл если все прошло успешно
		case false:
			shortUrl, err = u.GetUniqueString(longUrl)
			if err != nil {
				return "", errors.Wrap(err, "GetUniqueString")
			}

			if err := u.shortenerRepo.Create(ctx, longUrl, shortUrl); err != nil {
				return "", errors.Wrap(err, "u.shortenerRepo.CreateRepo")
			}
			if err := u.memdb.Create(shortUrl, longUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}

			logger.Infof("Return short url - %s", shortUrl)
			return shortUrl, nil
		}
	}
	logger.Infof("Return short url - %s", shortUrl)
	return shortUrl, nil
}

// Get return long url
func (u *USUseCase) Get(ctx context.Context, shortUrl string) (string, error) {
	if shortUrl == "" {
		logger.Error("Empty string in UrlShortenerUC.get")
		return "Your url is empty", nil
	}

	logger.Infof("New getting with short url===============>%s", shortUrl)

	// Проверяем есть ли урл локально
	longUrl, err := u.memdb.GetLong(shortUrl) 
	if err != nil {
		return "", errors.Wrap(err, "memdb.GetLong")
	}

	// Если локально длинный урл не найден ищем в базе
	if longUrl == "" { 
		urls, ok, err := u.shortenerRepo.Get(ctx, longUrl, shortUrl)
		if err != nil {
			return "", errors.Wrap(err, "shortenerRepo.Get")
		}

		switch ok {
		case true:
			if err := u.memdb.Create(urls.ShortUrl, urls.LongUrl); err != nil {
				return "", errors.Wrap(err, "u.memdb.CreateInmemory")
			}

			logger.Infof("Return long url - %s", urls.LongUrl)
			return urls.LongUrl, nil
		case false:
			return "This short url is not exist", nil
		}

		// longUrl = urls.LongUrl
	}
	logger.Infof("Return long url - %s", longUrl)
	return longUrl, nil
}

// getUniqueString create unique short url
func (u *USUseCase) GetUniqueString(longUrl string) (string, error) {
	hash := sha512.New()
	hash.Write([]byte(longUrl))
	x := hash.Sum([]byte("some salt here"))
	

	hd := hashids.NewData()

	hd.Salt = string(x)
	hd.Alphabet = u.cfg.Shortener.Runes
	hd.MinLength = u.cfg.Shortener.StringLength

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", errors.Wrap(err, "hashilds.NewWithData")
	}

	e, err := h.Encode([]int{434, 1313, 99})
	if err != nil {
		return "", errors.Wrap(err, "h.Encode")
	}

	return e, nil
}

func (u *USUseCase) CallAt(hour, min, sec int, sleepingDuration time.Duration, f func() error ) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()

	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		// Если получилось время раньше текущего, прибавляем сутки.
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	// Вычисляем временной промежуток до запуска.
	duration := firstCallTime.Sub(time.Now().Local())
	logger.Info("callat")
	go func() {
		time.Sleep(duration)
		for {
			f()
			// Следующий запуск через сутки.
			time.Sleep(u.cfg.CallAt.SleepingTime * sleepingDuration)
		}
	}()

	return nil
}


