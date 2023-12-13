package service_test

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"shrt-server/service"
	"shrt-server/test/mocks"
	"shrt-server/types"
	"shrt-server/types/entity"
	"testing"
)

func TestCreateShrtLink(t *testing.T) {
	t.Run("create with existing url", func(t *testing.T) {
		// Arrange
		control := gomock.NewController(t)
		defer control.Finish()

		expected := &types.CreateShortenLinkResponse{
			LongUrl: "https://www.google.com",
			Slug:    "abc",
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindByLongURL("https://www.google.com").Return(&entity.Shrt{
			ID:      1,
			LongURL: "https://www.google.com",
			Slug:    "abc",
		}, nil)
		shrtRepo.EXPECT().UpdateVisit(uint(1)).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		createUrl, _ := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})

		// Assert
		assert.Equal(t, expected, createUrl)
	})

	t.Run("create with existing url and check existing url error", func(t *testing.T) {
		// Arrange
		control := gomock.NewController(t)
		defer control.Finish()

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindByLongURL("https://www.google.com").Return(nil, types.ErrCheckExistingUrl)
		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})

		// Assert
		assert.Equal(t, types.ErrCheckExistingUrl, err)
	})

	t.Run("create with existing url but cannot update visit", func(t *testing.T) {
		// Arrange
		control := gomock.NewController(t)
		defer control.Finish()

		expected := &types.CreateShortenLinkResponse{
			LongUrl: "https://www.google.com",
			Slug:    "abc",
		}

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindByLongURL("https://www.google.com").Return(&entity.Shrt{
			ID:      1,
			LongURL: "https://www.google.com",
			Slug:    "abc",
		}, nil)

		shrtRepo.EXPECT().UpdateVisit(uint(1)).Return(types.ErrCannotUpdateVisit)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		createUrl, _ := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})

		// Assert
		assert.Equal(t, expected, createUrl)
	})

	invalidSlug := []string{"abc@", "ewdf*0", "e0il243+", "&3,123d", "rtrt12312=3", "34253!45", "#@$%%G"}

	for _, slug := range invalidSlug {
		t.Run("create link with invalid slug", func(t *testing.T) {
			// Arrange
			control := gomock.NewController(t)
			defer control.Finish()

			shrtRepo := mocks.NewMockShrtRepository(control)
			shrtRepo.EXPECT().FindByLongURL("https://www.google.com").Return(nil, gorm.ErrRecordNotFound)

			shrtService := service.NewShrtService(shrtRepo)

			// Act
			_, err := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
				LongURL: "https://www.google.com",
				Slug:    &slug, // invalid slug
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
			control := gomock.NewController(t)
			defer control.Finish()

			shrtRepo := mocks.NewMockShrtRepository(control)
			shrtRepo.EXPECT().FindByLongURL("https://www.google.com").Return(nil, gorm.ErrRecordNotFound)
			shrtRepo.EXPECT().FindBySlug(duplicatedSlug.slug).Return(&entity.Shrt{
				ID:      duplicatedSlug.id,
				LongURL: duplicatedSlug.longURL,
				Slug:    duplicatedSlug.slug,
			}, nil)

			shrtService := service.NewShrtService(shrtRepo)

			// Act
			_, err := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
				LongURL: "https://www.google.com",
				Slug:    &duplicatedSlug.slug,
			})

			// Assert
			assert.Equal(t, types.ErrSlugAlreadyExists, err)
		})
	}

	//t.Run("create link with slug", func(t *testing.T) {
	//
	//})
}

func TestGetOriginalUrl(t *testing.T) {
	t.Run("get original url with existing slug", func(t *testing.T) {
		// Arrange
		control := gomock.NewController(t)
		defer control.Finish()

		longUrl := "https://www.google.com"
		slug := "abcdefg"

		expected := &types.CreateShortenLinkResponse{
			LongUrl: longUrl,
			Slug:    slug,
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
		control := gomock.NewController(t)
		defer control.Finish()

		slug := "abcdefg"

		shrtRepo := mocks.NewMockShrtRepository(control)
		shrtRepo.EXPECT().FindBySlug(slug).Return(nil, gorm.ErrRecordNotFound)

		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.GetOriginalURL(slug)

		// Assert
		assert.Equal(t, types.ErrSlugNotFound, err)
	})
}

func TestGetOriginalUrlToRedirect(t *testing.T) {
	t.Run("get original url to redirect with existing slug", func(t *testing.T) {
		// Arrange
		control := gomock.NewController(t)
		defer control.Finish()

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
		control := gomock.NewController(t)
		defer control.Finish()

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
		control := gomock.NewController(t)
		defer control.Finish()

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
}
