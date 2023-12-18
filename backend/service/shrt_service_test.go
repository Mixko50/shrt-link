package service_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"shrt-server/service"
	"shrt-server/test/mocks"
	"shrt-server/types"
	"shrt-server/types/entity"
	"shrt-server/types/request"
	"shrt-server/types/response"
	"shrt-server/utilities/text"
	"testing"
)

func TestCreateShrtLink(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	invalidSlug := []string{"abc@", "ewdf*0", "e0il243+", "&3,123d", "rtrt12312=3", "34253!45", "#@$%%G"}

	for _, slug := range invalidSlug {
		t.Run("create link with invalid slug", func(t *testing.T) {
			// Arrange
			shrtRepo := mocks.NewMockShrtRepository(control)
			shrtService := service.NewShrtService(shrtRepo)

			// Act
			_, err := shrtService.CreateShrtLink(&request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        &slug, // invalid slug
			})

			// Assert
			assert.Equal(t, types.ErrSlugNotAlphanumeric, err)
		})
	}

	type duplicatedSlug struct {
		id      uint
		longURL string
		slug    string
	}

	duplicatedSlugList := []duplicatedSlug{
		{
			id:      1,
			longURL: "https://www.google.com",
			slug:    "abc",
		},
		{
			id:      2,
			longURL: "https://www.google.com",
			slug:    "google_test",
		},
		{
			id:      3,
			longURL: "https://www.google.com",
			slug:    "google_test_1",
		},
		{
			id:      4,
			longURL: "https://www.mixkomii.com",
			slug:    "mixko",
		},
	}

	for _, duplicatedSlug := range duplicatedSlugList {
		t.Run("create link with duplicated slug", func(t *testing.T) {
			// Arrange
			shrtRepo := mocks.NewMockShrtRepository(control)
			shrtRepo.EXPECT().FindBySlug(duplicatedSlug.slug).Return(&entity.Shrt{
				ID:      duplicatedSlug.id,
				LongURL: duplicatedSlug.longURL,
				Slug:    duplicatedSlug.slug,
			}, nil)

			shrtService := service.NewShrtService(shrtRepo)

			// Act
			_, err := shrtService.CreateShrtLink(&request.CreateShortenLinkRequest{
				OriginalUrl: "https://www.google.com",
				Slug:        &duplicatedSlug.slug,
			})

			// Assert
			assert.Equal(t, types.ErrSlugAlreadyExists, err)
		})
	}

	t.Run("create link with slug", func(t *testing.T) {
		// Arrange
		data := &request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("mixko_google"),
		}

		expected := &response.CreateShortenLinkResponse{
			OriginalUrl: "https://www.google.com",
			Slug:        "mixko_google",
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(*data.Slug).Return(nil, gorm.ErrRecordNotFound)
		shrtRepo.EXPECT().Create(&entity.Shrt{
			LongURL: data.OriginalUrl,
			Slug:    *data.Slug,
		}).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		createUrl, _ := shrtService.CreateShrtLink(data)

		// Assert
		assert.Equal(t, expected, createUrl)
	})

	t.Run("create link with slug and check existing slug error", func(t *testing.T) {
		// Arrange
		data := &request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("mixko_google"),
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(*data.Slug).Return(nil, types.ErrCheckExistingUrl)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.CreateShrtLink(data)

		// Assert
		assert.Equal(t, types.ErrCheckExistingUrl, err)
	})

	t.Run("create link with slug and cannot create shrt link", func(t *testing.T) {
		// Arrange
		data := &request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
			Slug:        text.Ptr("mixko_google"),
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(*data.Slug).Return(nil, gorm.ErrRecordNotFound)
		shrtRepo.EXPECT().Create(&entity.Shrt{
			LongURL: data.OriginalUrl,
			Slug:    *data.Slug,
		}).Return(types.ErrCannotCreateShrtLink)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.CreateShrtLink(data)

		// Assert
		assert.Equal(t, types.ErrCannotCreateShrtLink, err)
	})

	t.Run("create link without slug", func(t *testing.T) {
		// Arrange
		data := &request.CreateShortenLinkRequest{
			OriginalUrl: "https://www.google.com",
		}

		expectedLongUrl := "https://www.google.com"

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		shrtRepo := mocks.NewMockShrtRepository(ctrl)
		shrtRepo.EXPECT().Create(gomock.Any()).Do(func(input *entity.Shrt) {
			assert.NotNil(t, input.Slug) // Assert the generated slug
		}).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		createUrl, err := shrtService.CreateShrtLink(data)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expectedLongUrl, createUrl.OriginalUrl)
	})
}

func TestGetOriginalUrl(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	t.Run("get original url with existing slug", func(t *testing.T) {
		// Arrange
		longUrl := "https://www.google.com"
		slug := "abcdefg"

		expected := &response.CreateShortenLinkResponse{
			OriginalUrl: longUrl,
			Slug:        slug,
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(&entity.Shrt{
			ID:      1,
			LongURL: longUrl,
			Slug:    slug,
		}, nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		originalUrl, _ := shrtService.GetOriginalURL(slug)

		// Assert
		assert.Equal(t, expected, originalUrl)
	})

	t.Run("get original url with not existing slug", func(t *testing.T) {
		// Arrange
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(nil, gorm.ErrRecordNotFound)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.GetOriginalURL(slug)

		// Assert
		assert.Equal(t, types.ErrSlugNotFound, err)
	})

	t.Run("get original url with existing slug but cannot update visit", func(t *testing.T) {
		// Arrange
		longUrl := "https://www.google.com"
		slug := "abcdefg"

		expected := &response.CreateShortenLinkResponse{
			OriginalUrl: longUrl,
			Slug:        slug,
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(&entity.Shrt{
			ID:      1,
			LongURL: longUrl,
			Slug:    slug,
		}, nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		originalUrl, _ := shrtService.GetOriginalURL(slug)

		// Assert
		assert.Equal(t, expected, originalUrl)
	})

	t.Run("get original url with database error", func(t *testing.T) {
		// Arrange
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(nil, gorm.ErrInvalidTransaction)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.GetOriginalURL(slug)

		// Assert
		assert.Equal(t, types.ErrSomethingWentWrong, err)
	})
}

func TestGetOriginalUrlToRedirect(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	t.Run("get original url to redirect with existing slug", func(t *testing.T) {
		// Arrange
		longUrl := "https://www.google.com"
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(&entity.Shrt{
			ID:      1,
			LongURL: longUrl,
			Slug:    slug,
		}, nil)
		shrtRepo.EXPECT().UpdateVisit(uint(1)).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		originalUrl, _ := shrtService.GetOriginalURLToRedirect(slug)

		// Assert
		assert.Equal(t, longUrl, originalUrl)
	})

	t.Run("get original url to redirect with not existing slug", func(t *testing.T) {
		// Arrange
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(nil, gorm.ErrRecordNotFound)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.GetOriginalURLToRedirect(slug)

		// Assert
		assert.Equal(t, types.ErrSlugNotFound, err)
	})

	t.Run("get original url to redirect with existing slug but cannot update visit", func(t *testing.T) {
		// Arrange
		longUrl := "https://www.google.com"
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(&entity.Shrt{
			ID:      1,
			LongURL: longUrl,
			Slug:    slug,
		}, nil)
		shrtRepo.EXPECT().UpdateVisit(uint(1)).Return(types.ErrCannotUpdateVisit)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		originalUrl, _ := shrtService.GetOriginalURLToRedirect(slug)

		// Assert
		assert.Equal(t, longUrl, originalUrl)
	})

	t.Run("get original url to redirect with database error", func(t *testing.T) {
		// Arrange
		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(nil, gorm.ErrInvalidTransaction)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.GetOriginalURLToRedirect(slug)

		// Assert
		assert.Equal(t, types.ErrSomethingWentWrong, err)
	})
}
