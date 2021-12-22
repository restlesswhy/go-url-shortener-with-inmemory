package inmemory

import (
	"time"

	"github.com/hashicorp/go-memdb"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

type Urls struct {
	ShortUrl string
	LongUrl string
	ExpiresAt string
}

type USInmemory struct {
	memdb *memdb.MemDB
}

func NewUSInmemory(memdb *memdb.MemDB) *USInmemory {
	return &USInmemory{
		memdb: memdb,
	}
}

// CreateInmemory create url in local memory
func (u *USInmemory) Create(shortUrl, longUrl string) error {
	logger.Info("Creating new url in cache")

	layout := "2006-01-02 15:04:05"

	txn := u.memdb.Txn(true)
	urls := &Urls{
		ShortUrl: shortUrl,
		LongUrl: longUrl,
		ExpiresAt: time.Now().Add(1 * time.Hour).Format(layout),
	}
	
	if err := txn.Insert("urls", urls); err != nil {
		return errors.Wrap(err, "Create.Insert")
	}
	txn.Commit()

	logger.Infof("Added in cache: longUrl - %s, shortUrl - %s", longUrl, shortUrl)
	return nil
}

// GetShortInmemory return short url if exist
func (u *USInmemory) GetShort(longUrl string) (string, error) {
	defer func() {
        if err := recover(); err != nil {
			logger.Info("Have no in cache")
        }
    }()

	logger.Info("Serching in cache...")

	txn := u.memdb.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("urls", "longUrl", longUrl)
	if err != nil {
		return "", errors.Wrap(err, "GetShort.First")
	}
	
	shortUrl := raw.(*Urls).ShortUrl

	logger.Infof("Found url in cache")
	return shortUrl, nil
}

// GetLongInmemory return long url if exist
func (u *USInmemory) GetLong(shortUrl string) (string, error) {
	defer func() {
        if err := recover(); err != nil {
			logger.Info("Have no in cache")
        }
    }()

	logger.Info("Serching in cache...")

	txn := u.memdb.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("urls", "id", shortUrl)
	if err != nil {
		return "", errors.Wrap(err, "GetLong.First")
	}

	longUrl := raw.(*Urls).LongUrl

	logger.Infof("Found url in cache")
	return longUrl, nil
}

func (u *USInmemory) Check() error {
	txn := u.memdb.Txn(false)

	logger.Info("urls are checking..")
	it, err := txn.Get("urls", "id")
	if err != nil {
		logger.Error("Check.Get", err)
	}
	txn.Abort()

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Urls)

		expiresAt := parseTime(p.ExpiresAt)
		logger.Info(expiresAt, p.ShortUrl, time.Now())
		
		if !expiresAt.After(time.Now()) {
			logger.Infof("deleted url - %s", p.LongUrl)
			if err := u.Delete(p.ShortUrl); err != nil {
				return errors.Wrap(err, "Check.Delete")
			}
		}
		
	}	
	logger.Info("urls were checked")
	return nil
}

func (u *USInmemory) Delete(shortUrl string) error {
	txn := u.memdb.Txn(true)
	defer txn.Abort()

	if _, err := txn.DeleteAll("urls", "id", shortUrl); err != nil {
		return errors.Wrap(err, "txn.DeleteAll")
	}
	txn.Commit()
	return nil
}

func parseTime(timeStr string) time.Time {
	layout := "2006-01-02 15:04:05"

	ti, err := time.Parse(layout, timeStr)
	if err != nil {
		logger.Error("pareseTime", err)
	}
	
	return ti
}