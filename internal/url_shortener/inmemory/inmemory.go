package inmemory

import (

	"github.com/hashicorp/go-memdb"
	"github.com/restlesswhy/grpc/url-shortener-microservice/pkg/logger"
)

type Urls struct {
	ShortUrl string
	LongUrl string
}

type UrlShortenerInmemory struct {
	memdb *memdb.MemDB
}

func NewUrlShortenerInmemory(memdb *memdb.MemDB) *UrlShortenerInmemory {
	return &UrlShortenerInmemory{
		memdb: memdb,
	}
}


func (u *UrlShortenerInmemory) CreateInmemory(shortUrl, longUrl string) error {
	logger.Info("creating new url in inmemory")
	logger.Info(shortUrl, longUrl)
	var err error

	txn := u.memdb.Txn(true)
	urls := &Urls{
		ShortUrl: shortUrl,
		LongUrl: longUrl,
	}
	
	if err = txn.Insert("urls", urls); err != nil {
		return err
	}

	txn.Commit()

	logger.Infof("added to memdb: shortUrl - %s, longUrl - %s", shortUrl, longUrl)
	return nil
}

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

func (u *UrlShortenerInmemory) GetLongInmemory(shortUrl string) (string, error) {
	defer func() {
        if err := recover(); err != nil {
			logger.Info("have no in memdb")
        }
    }()

	logger.Info("serching in inmemory")
	txn := u.memdb.Txn(false)
	defer txn.Abort()
	// txn.First()
	raw, _ := txn.First("urls", "id", shortUrl)
	
	longUrl := raw.(*Urls).LongUrl
	logger.Infof("found url in inmemory - %s", longUrl)
	return longUrl, nil
}