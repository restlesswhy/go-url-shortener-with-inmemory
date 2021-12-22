package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/restlesswhy/grpc/url-shortener-microservice/config"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/models"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/mock"
	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"github.com/stretchr/testify/require"
)



func TestUrlShortenerUC_CreateEmptyErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(nil, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UCRequest{
		LongUrl: "",
	}
	
	shortUrlExpect := "Your url is empty"

	shortUrl, err := shortUC.Create(ctx, reqValue.LongUrl)

	require.Nil(t, err)
	require.Equal(t, shortUrlExpect, shortUrl)
}

func TestUrlShortenerUC_CreateInmemory(t *testing.T) {
	// Если найден в кэше
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	cfg := &config.Config{
		Shortener: config.ShortenerConfig{
			StringLength: 10,
			Runes: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_",
		},
	}

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(cfg, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UCRequest{
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}
	
	expectShortUrl, _ := shortUC.GetUniqueString(reqValue.LongUrl)

	shortMemDB.EXPECT().GetShort(gomock.Any()).Return(expectShortUrl, nil)

	shortUrl, err := shortUC.Create(ctx, reqValue.LongUrl)

	require.Nil(t, err)
	require.Equal(t, expectShortUrl, shortUrl)
}

func TestUrlShortenerUC_CreateFindInRepo(t *testing.T) {
	// Если найден в базе
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	cfg := &config.Config{
		Shortener: config.ShortenerConfig{
			StringLength: 10,
			Runes: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_",
		},
	}

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(cfg, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UCRequest{
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}
	

	shortMemDB.EXPECT().GetShort(gomock.Any()).Return("", nil)
	expectShortUrl, _ := shortUC.GetUniqueString(reqValue.LongUrl)

	model :=  models.UrlsLS{
		ShortUrl: expectShortUrl,
	}

	shortRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(model, true, nil)
	shortMemDB.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	shortUrl, err := shortUC.Create(ctx, reqValue.LongUrl)

	require.Nil(t, err)
	require.Equal(t, expectShortUrl, shortUrl)
}

func TestUrlShortenerUC_Create(t *testing.T) {
	// Если нигде не найден
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &config.Config{
		Shortener: config.ShortenerConfig{
			StringLength: 10,
			Runes: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_",
		},
	}

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(cfg, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UCRequest{
		LongUrl: "https://habr.com/ru/company/vk/blog/476276/",
	}

	shortMemDB.EXPECT().GetShort(gomock.Any()).Return("", nil)
	expectShortUrl, _ := shortUC.GetUniqueString(reqValue.LongUrl)

	model :=  models.UrlsLS{
		ShortUrl: expectShortUrl,
	}

	shortRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(model, false, nil)
	shortRepo.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	shortMemDB.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()

	shortUrl, err := shortUC.Create(ctx, reqValue.LongUrl)

	require.Nil(t, err)
	require.Equal(t, expectShortUrl, shortUrl)
}

func TestUrlShortenerUC_GetEmptyErr(t *testing.T) {
	// Если передана пустая строка
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(nil, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UGRequest{
		ShortUrl: "",
	}
	
	longUrlExpect := "Your url is empty"

	ctx := context.Background()
	longUrl, err := shortUC.Get(ctx, reqValue.ShortUrl)

	require.Nil(t, err)
	require.Equal(t, longUrlExpect, longUrl)
}

func TestUrlShortenerUC_GetInmemory(t *testing.T) {
	// Если найден в кэше
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(nil, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UGRequest{
		ShortUrl: "some short url",
	}
	
	expectLongUrl := "some long url"

	shortMemDB.EXPECT().GetLong(gomock.Any()).Return(expectLongUrl, nil)

	longUrl, err := shortUC.Get(ctx, reqValue.ShortUrl)

	require.Nil(t, err)
	require.Equal(t, expectLongUrl, longUrl)
}

func TestUrlShortenerUC_GetBase(t *testing.T) {
	// Если найден в базе
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(nil, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UGRequest{
		ShortUrl: "some short url",
	}
	
	expectLongUrl := "some long url"

	model :=  models.UrlsLS{
		LongUrl: expectLongUrl,
	}

	shortMemDB.EXPECT().GetLong(gomock.Any()).Return("", nil)
	shortRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(model, true, nil)
	shortMemDB.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	longUrl, err := shortUC.Get(ctx, reqValue.ShortUrl)

	require.Nil(t, err)
	require.Equal(t, expectLongUrl, longUrl)
}

func TestUrlShortenerUC_Get(t *testing.T) {
	// Если не найден нигде
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	ctx := context.Background()

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(nil, shortRepo, shortMemDB)
	
	reqValue := &shortenerService.UGRequest{
		ShortUrl: "some short url",
	}
	
	expectLongUrl := "This short url is not exist"

	model :=  models.UrlsLS{
		LongUrl: expectLongUrl,
	}

	shortMemDB.EXPECT().GetLong(gomock.Any()).Return("", nil)
	shortRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(model, false, nil)

	longUrl, err := shortUC.Get(ctx, reqValue.ShortUrl)

	require.Nil(t, err)
	require.Equal(t, expectLongUrl, longUrl)
}

func TestUrlShortenerUC_CallAt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	cfg := &config.Config{
		CallAt: config.CallAtConfig{
			SleepingTime: 2,
		},
	}

	shortRepo := mock.NewMockUSRepository(ctrl)
	shortMemDB := mock.NewMockUSInmemory(ctrl)
	shortUC := NewUSUseCase(cfg, shortRepo, shortMemDB)

	x := 0

	err := shortUC.CallAt(time.Now().Hour(), time.Now().Minute(), time.Now().Add(2 * time.Second).Second(), time.Second, func() error {
		x++
		return nil
	})
	
	time.Sleep(10 * time.Second)
	xExpect := 5

	require.Nil(t, err)
	require.Equal(t, xExpect, x)
}