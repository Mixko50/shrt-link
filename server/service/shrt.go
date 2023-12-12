package service

import (
	"shrt-server/types"
)

type ShrtService interface {
	CreateShrtLink(body *types.CreateShortenLinkRequest) (*types.CreateShortenLinkResponse, error)
	GetOriginalURL(slug string) (*types.CreateShortenLinkResponse, error)
	GetOriginalURLToRedirect(slug string) (string, error)
}
