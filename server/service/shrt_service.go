package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shrt-server/repository"
	"shrt-server/types"
	"shrt-server/types/entity"
	"shrt-server/types/request"
	"shrt-server/types/response"
	"shrt-server/utilities/text"
)

type shrtService struct {
	shrtRepository repository.ShrtRepository
}

func NewShrtService(shrtRepository repository.ShrtRepository) shrtService {
	return shrtService{shrtRepository: shrtRepository}
}

func (s shrtService) CreateShrtLink(body *request.CreateShortenLinkRequest) (*response.CreateShortenLinkResponse, error) {
	var slug string
	if body.Slug != nil {
		// validate slug
		if !text.IsAlphanumeric(*body.Slug) {
			return nil, types.ErrSlugNotAlphanumeric
		}

		// check duplicated slug
		checkDuplicated, err := s.shrtRepository.FindBySlug(*body.Slug)
		if !errors.Is(err, gorm.ErrRecordNotFound) && checkDuplicated != nil {
			return nil, types.ErrSlugAlreadyExists
		}

		// other error
		if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
			return nil, types.ErrCheckExistingUrl
		}

		slug = *body.Slug
	} else {
		slug = text.GenerateSlug()
	}

	if err := s.shrtRepository.Create(&entity.Shrt{
		LongURL: body.OriginalUrl,
		Slug:    slug,
	}); err != nil {
		return nil, types.ErrCannotCreateShrtLink
	}

	return &response.CreateShortenLinkResponse{
		OriginalUrl: body.OriginalUrl,
		Slug:        slug,
	}, nil
}

func (s shrtService) GetOriginalURL(slug string) (*response.CreateShortenLinkResponse, error) {
	shortLink, err := s.shrtRepository.FindBySlug(slug)
	if err != nil {
		return nil, types.ErrSlugNotFound
	}

	return &response.CreateShortenLinkResponse{
		OriginalUrl: shortLink.LongURL,
		Slug:        shortLink.Slug,
	}, nil
}

func (s shrtService) GetOriginalURLToRedirect(slug string) (string, error) {
	shortLink, err := s.shrtRepository.FindBySlug(slug)
	if err != nil {
		return "", types.ErrSlugNotFound
	}

	err = s.shrtRepository.UpdateVisit(shortLink.ID)
	if err != nil {
		log.Error(err)
	}

	return shortLink.LongURL, nil
}
