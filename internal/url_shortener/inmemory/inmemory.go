package inmemory

import (
	"time"

	"github.com/hashicorp/go-memdb"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

type Urls struct {
	ShortUrl string
	LongUrl string
	ExpiresAt string
}

type UrlShortenerInmemory struct {
	memdb *memdb.MemDB
}

func NewUrlShortenerInmemory(memdb *memdb.MemDB) *UrlShortenerInmemory {
	return &UrlShortenerInmemory{
		memdb: memdb,
	}
}

// CreateInmemory create url in local memory
func (u *UrlShortenerInmemory) CreateInmemory(shortUrl, longUrl string) error {
	logger.Info("creating new url in inmemory")

	var err error
	layout := "2006-01-02 15:04:05"

	txn := u.memdb.Txn(true)
	urls := &Urls{
		ShortUrl: shortUrl,
		LongUrl: longUrl,
		ExpiresAt: time.Now().Add(1 * time.Hour).Format(layout),
	}
	
	if err = txn.Insert("urls", urls); err != nil {
		return err
	}

	txn.Commit()


	logger.Infof("added to memdb: longUrl - %s, shortUrl - %s", longUrl, shortUrl)
	return nil
}

// GetShortInmemory return short url if exist
func (u *UrlShortenerInmemory) GetShortInmemory(longUrl string) (string, error) {
	defer func() {
        if err := recover(); err != nil {
			logger.Info("have no in memdb")
        }
    }()

	logger.Info("serching in inmemory")
	txn := u.memdb.Txn(false)
	defer txn.Abort()
	// txn.First()
	raw, _ := txn.First("urls", "longUrl", longUrl)
	
	shortUrl := raw.(*Urls).ShortUrl
	logger.Infof("found url in inmemory - %s", shortUrl)
	return shortUrl, nil
}

// GetLongInmemory return long url if exist
func (u *UrlShortenerInmemory) GetLongInmemory(shortUrl string) (string, error) {
	defer func() {
        if err := recover(); err != nil {
			logger.Info("have no in memdb")
        }
    }()

	logger.Info("serching in inmemory")

	txn := u.memdb.Txn(false)
	defer txn.Abort()

	raw, _ := txn.First("urls", "id", shortUrl)
	
	longUrl := raw.(*Urls).LongUrl
	logger.Infof("found url in inmemory - %s", longUrl)
	return longUrl, nil
}

func (u *UrlShortenerInmemory) CheckInmemory() {
	txn := u.memdb.Txn(false)

	logger.Info("lets check")
	it, err := txn.Get("urls", "id")
	if err != nil {
		logger.Error("CheckInmemory.Get", err)
	}
	txn.Abort()

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Urls)

		expiresAt := parseTime(p.ExpiresAt)
		logger.Info(expiresAt, p.ShortUrl, time.Now())
		
		if !expiresAt.After(time.Now()) {
			logger.Info("time 2 del")
			u.Delete(p.ShortUrl)
		}
		
	}	
	logger.Info("checked")
}

func (u *UrlShortenerInmemory) Delete(shortUrl string) {
	txn := u.memdb.Txn(true)
	defer txn.Abort()

	if _, err := txn.DeleteAll("urls", "id", shortUrl); err != nil {
		logger.Error("CheckInmemory.Delete", err)
	}
	txn.Commit()
}

func parseTime(timeStr string) time.Time {
	layout := "2006-01-02 15:04:05"

	ti, err := time.Parse(layout, timeStr)
	if err != nil {
		logger.Error("pareseTime", err)
	}
	
    
	return ti
}