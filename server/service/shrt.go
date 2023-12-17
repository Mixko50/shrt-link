package service

import (
	"shrt-server/types/request"
	"shrt-server/types/response"
)

type ShrtService interface {
	CreateShrtLink(body *request.CreateShortenLinkRequest) (*response.CreateShortenLinkResponse, error)
	GetOriginalURL(slug string) (*response.CreateShortenLinkResponse, error)
	GetOriginalURLToRedirect(slug string) (string, error)
}
