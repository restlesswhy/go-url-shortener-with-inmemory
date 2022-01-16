package grpcdel

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/mock"
	shortenerService "github.com/restlesswhy/grpc/url-shortener-microservice/internal/url_shortener/proto"
	"github.com/stretchr/testify/require"
)

func TestUrlShortenerMicroservice_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	shortUC := mock.NewMockUSUseCase(ctrl)
	shortServerGRPC := NewUSMicroservice(shortUC)

	reqValue := &shortenerService.UCRequest{
		LongUrl: "https://translate.google.com/?hl=ru&sl=en&tl=ru&text=Parallel%20signals%20that%20this%20test%20is%20to%20be%20run%20in",
	}
	
	t.Run("Create", func(t *testing.T) {

		shortUrl := "bWY4CqPTMg"
		

		shortUC.EXPECT().Create(gomock.Any(), gomock.Any()).Return(shortUrl, nil)
		
		response, err := shortServerGRPC.Create(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, shortUrl, response.ShortUrl)
	})
}

func TestUrlShortenerMicroservice_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	shortUC := mock.NewMockUSUseCase(ctrl)
	shortServerGRPC := NewUSMicroservice(shortUC)

	reqValue := &shortenerService.UGRequest{
		ShortUrl: "bWY4CqPTMg",
	}
	
	t.Run("Create", func(t *testing.T) {

		longUrl := "https://translate.google.com/?hl=ru&sl=en&tl=ru&text=Parallel%20signals%20that%20this%20test%20is%20to%20be%20run%20in"
		

		shortUC.EXPECT().Get(gomock.Any(), gomock.Any()).Return(longUrl, nil)
		
		response, err := shortServerGRPC.Get(context.Background(), reqValue)
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Equal(t, longUrl, response.LongUrl)
	})
}
