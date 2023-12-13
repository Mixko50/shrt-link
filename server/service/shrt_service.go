package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shrt-server/repository"
	"shrt-server/types"
	"shrt-server/types/entity"
	"shrt-server/utilities/text"
)

type shrtService struct {
	shrtRepository repository.ShrtRepository
}

func NewShrtService(shrtRepository repository.ShrtRepository) shrtService {
	return shrtService{shrtRepository: shrtRepository}
}

func (s shrtService) CreateShrtLink(body *types.CreateShortenLinkRequest) (*types.CreateShortenLinkResponse, error) {
	shrt, err := s.shrtRepository.FindByLongURL(body.LongURL)
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, types.ErrCheckExistingUrl
	}

	if shrt != nil {
		err := s.shrtRepository.UpdateVisit(shrt.ID)
		if err != nil {
			log.Error(types.ErrCannotUpdateVisit)
		}

		return &types.CreateShortenLinkResponse{
			LongUrl: shrt.LongURL,
			Slug:    shrt.Slug,
		}, nil
	}

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
		slug = *body.Slug
	} else {
		slug = text.GenerateSlug()
	}

	if err := s.shrtRepository.Create(&entity.Shrt{
		LongURL: body.LongURL,
		Slug:    slug,
	}); err != nil {
		return nil, types.ErrCannotCreateShrtLink
	}

	return &types.CreateShortenLinkResponse{
		LongUrl: body.LongURL,
		Slug:    slug,
	}, nil
}

func (s shrtService) GetOriginalURL(slug string) (*types.CreateShortenLinkResponse, error) {
	shortLink, err := s.shrtRepository.FindBySlug(slug)
	if err != nil {
		return nil, types.ErrSlugNotFound
	}

	return &types.CreateShortenLinkResponse{
		LongUrl: shortLink.LongURL,
		Slug:    shortLink.Slug,
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
