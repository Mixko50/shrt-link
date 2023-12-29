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
	"shrt-server/utility"
	"shrt-server/utility/text"
)

type shrtService struct {
	shrtRepository repository.ShrtRepository
	utility        utility.Utility
}

func NewShrtService(shrtRepository repository.ShrtRepository, utlity utility.Utility) shrtService {
	return shrtService{
		shrtRepository: shrtRepository,
		utility:        utlity,
	}
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
		slug = s.utility.GenerateSlug()
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
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, types.ErrSlugNotFound
	}

	if err != nil {
		return nil, types.ErrSomethingWentWrong
	}

	return &response.CreateShortenLinkResponse{
		OriginalUrl: shortLink.LongURL,
		Slug:        shortLink.Slug,
	}, nil
}

func (s shrtService) GetOriginalURLToRedirect(slug string) (string, error) {
	shortLink, err := s.shrtRepository.FindBySlug(slug)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return "", types.ErrSlugNotFound
	}

	if err != nil {
		return "", types.ErrSomethingWentWrong
	}

	err = s.shrtRepository.UpdateVisit(shortLink.ID)
	if err != nil {
		log.Error(err)
	}

	return shortLink.LongURL, nil
}
