package service_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"shrt-server/repository"
	"shrt-server/service"
	"shrt-server/types"
	"shrt-server/types/entity"
	"shrt-server/utilities/text"
	"testing"
)

func TestCreateShrtLink(t *testing.T) {
	t.Run("create with existing url", func(t *testing.T) {
		// Arrange
		shrtRepo := repository.NewShrtRepositoryMock()
		shrtRepo.On("FindByLongURL", "https://www.google.com").Return(&entity.Shrt{
			ID:      1,
			LongURL: "https://www.google.com",
			Slug:    "abc",
		}, nil)

		shrtRepo.On("UpdateVisit", uint(1)).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)
		// Act

		createUrl, _ := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})
		expected := &types.CreateShortenLinkResponse{
			LongUrl: "https://www.google.com",
			Slug:    "abc",
		}

		// Assert
		assert.Equal(t, expected, createUrl)
	})

	t.Run("create with existing url and check existing url error", func(t *testing.T) {
		// Arrange
		shrtRepo := repository.NewShrtRepositoryMock()
		shrtRepo.On("FindByLongURL", "https://www.google.com").Return(&entity.Shrt{}, types.ErrCheckExistingUrl)
		shrtService := service.NewShrtService(shrtRepo)

		// Act
		_, err := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})
		fmt.Println(err)

		// Assert
		assert.Equal(t, types.ErrCheckExistingUrl, err)
	})

	t.Run("create with existing url but cannot update visit", func(t *testing.T) {
		// Arrange
		shrtRepo := repository.NewShrtRepositoryMock()
		shrtRepo.On("FindByLongURL", "https://www.google.com").Return(&entity.Shrt{
			ID:      1,
			LongURL: "https://www.google.com",
			Slug:    "abc",
		}, nil)

		shrtRepo.On("UpdateVisit", uint(1)).Return(types.ErrCannotUpdateVisit)

		shrtService := service.NewShrtService(shrtRepo)
		// Act

		createUrl, _ := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
		})
		expected := &types.CreateShortenLinkResponse{
			LongUrl: "https://www.google.com",
			Slug:    "abc",
		}

		// Assert
		assert.Equal(t, expected, createUrl)
	})

	t.Run("create link with invalid slug", func(t *testing.T) {
		// Arrange
		shrtRepo := repository.NewShrtRepositoryMock()
		shrtRepo.On("FindByLongURL", "https://www.google.com").Return(nil, gorm.ErrRecordNotFound)

		shrtRepo.On("UpdateVisit", uint(1)).Return(nil)

		shrtService := service.NewShrtService(shrtRepo)
		// Act

		createUrl, _ := shrtService.CreateShrtLink(&types.CreateShortenLinkRequest{
			LongURL: "https://www.google.com",
			Slug:    text.Ptr("abc@"), // invalid slug
		})

		// Assert
		assert.Equal(t, types.ErrSlugNotAlphanumeric, createUrl)
	})
}
